package task

import (
	"context"
	"fmt"
	"github.com/influxdata/flux"
	"github.com/influxdata/influxdb/v2/kit/platform"
	"github.com/influxdata/influxdb/v2/kit/tracing"
	"github.com/influxdata/influxdb/v2/query"
	taskmodel "github.com/uvite/gvmdesk/pkg/model"
	"go.uber.org/zap"
	"sync"
	"time"
)

const (
	maxPromises       = 1000
	defaultMaxWorkers = 100

	lastSuccessOption = "tasks.lastSuccessTime"
)

// promise represents a promise the executor makes to finish a run's execution asynchronously.
type promise struct {
	run  *taskmodel.Run
	task *taskmodel.Task

	done chan struct{}
	err  error

	createdAt time.Time
	startedAt time.Time

	ctx        context.Context
	cancelFunc context.CancelFunc
}

// ID is the id of the run that was created
func (p *promise) ID() platform.ID {
	return p.run.ID
}

// Cancel is used to cancel a executing query
func (p *promise) Cancel(ctx context.Context) {
	// call cancelfunc
	p.cancelFunc()

	// wait for ctx.Done or p.Done
	select {
	case <-p.Done():
	case <-ctx.Done():
	}
}

// Done provides a channel that closes on completion of a promise
func (p *promise) Done() <-chan struct{} {
	return p.done
}

// Error returns the error resulting from a run execution.
// If the execution is not complete error waits on Done().
func (p *promise) Error() error {
	<-p.done
	return p.err
}

// Executor it a task specific executor that works with the new scheduler system.
type Executor struct {
	log *zap.Logger
	ts  taskmodel.TaskService
	qs  query.QueryService

	// currentPromises are all the promises we are made that have not been fulfilled
	currentPromises sync.Map

	// futurePromises are promises that are scheduled to be executed in the future
	futurePromises sync.Map

	// keep a pool of promise's we have in queue
	promiseQueue chan *promise

	// keep a pool of execution workers.
	workerPool  sync.Pool
	workerLimit chan struct{}
}

// NewExecutor creates a new task executor
func NewExecutor(log *zap.Logger, ts taskmodel.TaskService) *Executor {

	e := &Executor{
		log: log,
		ts:  ts,

		currentPromises: sync.Map{},
		futurePromises:  sync.Map{},
		promiseQueue:    make(chan *promise, maxPromises),
		workerLimit:     make(chan struct{}, 100),
	}

	wm := &workerMaker{
		e: e,
	}

	go e.processScheduledTasks()

	e.workerPool = sync.Pool{New: wm.new}
	return e
}

func (e *Executor) startWorker() {
	// see if have available workers
	select {
	case e.workerLimit <- struct{}{}:
	default:
		// we have reached our worker limit and we cannot start any more.
		return
	}
	// fire up some workers
	worker := e.workerPool.Get().(*worker)
	if worker != nil {
		// if the worker is nil all the workers are busy and one of them will pick up the work we enqueued.
		go func() {
			// don't forget to put the worker back when we are done
			defer e.workerPool.Put(worker)
			worker.work()

			// remove a struct from the worker limit to another worker to work
			<-e.workerLimit
		}()
	}
}

// Cancel a run of a specific task.
func (e *Executor) Cancel(ctx context.Context, runID platform.ID) error {
	// find the promise
	val, ok := e.currentPromises.Load(runID)
	if !ok {
		return nil
	}
	promise := val.(*promise)

	// call cancel on it.
	promise.Cancel(ctx)

	return nil
}

func (e *Executor) processScheduledTasks() {
	t := time.Tick(1 * time.Second)
	for range t {
		e.futurePromises.Range(func(k any, v any) bool {
			vv := v.(*promise)
			if vv.run.ScheduledFor.Equal(time.Now()) || vv.run.ScheduledFor.Before(time.Now()) {
				if vv.run.RunAt.IsZero() {
					e.promiseQueue <- vv
					e.futurePromises.Delete(k)
					e.startWorker()
				}
			}
			return true
		})
	}
}
func (e *Executor) LoadExistingScheduleRuns(ctx context.Context) error {
	tasks, _, err := e.ts.FindTasks(ctx, taskmodel.TaskFilter{})
	if err != nil {
		e.log.Error("err finding tasks:", zap.Error(err))
		return err
	}
	for _, t := range tasks {
		beforeTime := time.Now().Add(time.Hour * 24 * 365).Format(time.RFC3339)
		runs, _, err := e.ts.FindRuns(ctx, taskmodel.RunFilter{Task: t.ID, BeforeTime: beforeTime})
		if err != nil {
			e.log.Error("err finding runs:", zap.Error(err))
			return err
		}
		for _, run := range runs {
			if run.ScheduledFor.After(time.Now()) {

				ctx, cancel := context.WithCancel(ctx)
				// create promise
				p := &promise{
					run:        run,
					task:       t,
					createdAt:  time.Now().UTC(),
					done:       make(chan struct{}),
					ctx:        ctx,
					cancelFunc: cancel,
				}
				e.futurePromises.Store(run.ID, p)
			}
		}
	}

	return nil
}

type workerMaker struct {
	e *Executor
}

func (wm *workerMaker) new() interface{} {
	return &worker{
		e: wm.e,
	}
}

type worker struct {
	e *Executor

	// exhaustResultIterators is used to exhaust the result
	// of a flux query
	exhaustResultIterators func(res flux.Result) error
}

func (w *worker) work() {
	// loop until we have no more work to do in the promise queue
	for {
		var prom *promise
		// check to see if we can execute
		select {
		case p, ok := <-w.e.promiseQueue:

			if !ok {
				// the promiseQueue has been closed
				return
			}
			prom = p
		default:
			// if nothing is left in the queue we are done
			return
		}

		// check to make sure we are below the limits.
		for {

			// sleep
			select {
			// If done the promise was canceled
			case <-prom.ctx.Done():
				prom.err = taskmodel.ErrRunCanceled
				close(prom.done)
				return
			case <-time.After(time.Second):
			}
		}

		// execute the promise
		w.executeQuery(prom)

		// close promise done channel and set appropriate error
		close(prom.done)

		// remove promise from registry
		w.e.currentPromises.Delete(prom.run.ID)
	}
}

func (w *worker) start(p *promise) {
	// trace
	span, ctx := tracing.StartSpanFromContext(p.ctx)
	defer span.Finish()
	fmt.Println(ctx)

	p.startedAt = time.Now()
}

func (w *worker) finish(p *promise, rs taskmodel.RunStatus, err error) {
	span, ctx := tracing.StartSpanFromContext(p.ctx)
	defer span.Finish()
	fmt.Println(ctx)

}

func (w *worker) executeQuery(p *promise) {
	span, ctx := tracing.StartSpanFromContext(p.ctx)
	defer span.Finish()

	// start
	w.start(p)
	fmt.Println(ctx)

	w.finish(p, taskmodel.RunSuccess, nil)
}
