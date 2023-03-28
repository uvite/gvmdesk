package taskmodel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uvite/gvmdesk/pkg/platform"
	"strconv"
	"time"


	"github.com/influxdata/influxdb/v2/task/options"
)

const (
	TaskDefaultPageSize = 100
	TaskMaxPageSize     = 500

	// TODO(jsteenb2): make these constants of type Status

	TaskStatusActive   = "active"
	TaskStatusInactive = "inactive"
)

var (
	// TaskSystemType is the type set in tasks' for all crud requests
	TaskSystemType = "system"
	// TaskBasicType is short-hand used by the UI to request a minimal subset of system task metadata
	TaskBasicType = "basic"
)

// Task is a task. 🎊
type Task struct {
	ID             platform.ID `json:"id"`
	Type           string      `json:"type,omitempty"`
	OrganizationID platform.ID `json:"orgID"`
	Organization   string      `json:"org"`
	OwnerID        platform.ID `json:"ownerID"`

	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	Content         string                 `json:"content"`  // 内容
	Path            string                 `json:"path"`     // 内容
	Symbol          string                 `json:"symbol"`   // 标签
	Interval        string                 `json:"interval"` // 标签
	Status          string                 `json:"status"`
	Flux            string                 `json:"flux"`
	Every           string                 `json:"every,omitempty"`
	Cron            string                 `json:"cron,omitempty"`
	Offset          time.Duration          `json:"offset,omitempty"`
	LatestCompleted time.Time              `json:"latestCompleted,omitempty"`
	LatestScheduled time.Time              `json:"latestScheduled,omitempty"`
	LatestSuccess   time.Time              `json:"latestSuccess,omitempty"`
	LatestFailure   time.Time              `json:"latestFailure,omitempty"`
	LastRunStatus   string                 `json:"lastRunStatus,omitempty"`
	LastRunError    string                 `json:"lastRunError,omitempty"`
	CreatedAt       time.Time              `json:"createdAt,omitempty"`
	UpdatedAt       time.Time              `json:"updatedAt,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

// EffectiveCron returns the effective cron string of the options.
// If the cron option was specified, it is returned.
// If the every option was specified, it is converted into a cron string using "@every".
// Otherwise, the empty string is returned.
// The value of the offset option is not considered.
func (t *Task) EffectiveCron() string {
	if t.Cron != "" {
		return t.Cron
	}
	if t.Every != "" {
		return "@every " + t.Every
	}
	return ""
}

// Run is a record createId when a run of a task is scheduled.
type Run struct {
	ID           platform.ID `json:"id,omitempty"`
	TaskID       platform.ID `json:"taskID"`
	Status       string      `json:"status"`
	ScheduledFor time.Time   `json:"scheduledFor"`          // ScheduledFor is the Now time used in the task's query
	RunAt        time.Time   `json:"runAt"`                 // RunAt is the time the task is scheduled to be run, which is ScheduledFor + Offset
	Flux         string      `json:"flux"`                  // Flux used in run
	StartedAt    time.Time   `json:"startedAt,omitempty"`   // StartedAt is the time the executor begins running the task
	FinishedAt   time.Time   `json:"finishedAt,omitempty"`  // FinishedAt is the time the executor finishes running the task
	RequestedAt  time.Time   `json:"requestedAt,omitempty"` // RequestedAt is the time the coordinator told the scheduler to schedule the task
	Log          []Log       `json:"log,omitempty"`

	TraceID   string `json:"traceID"`   // TraceID preserves the trace id
	IsSampled bool   `json:"isSampled"` // IsSampled preserves whether this run was sampled
}

// Log represents a link to a log resource
type Log struct {
	RunID   platform.ID `json:"runID,omitempty"`
	Time    string      `json:"time"`
	Message string      `json:"message"`
}

func (l Log) String() string {
	return l.Time + ": " + l.Message
}

// TaskService represents a service for managing one-off and recurring tasks.
type TaskService interface {
	// FindTaskByID returns a single task
	FindTaskByID(ctx context.Context, id platform.ID) (*Task, error)

	// FindTasks returns a list of tasks that match a filter (limit 100) and the total count
	// of matching tasks.
	FindTasks(ctx context.Context, filter TaskFilter) ([]*Task, int, error)

	// CreateTask creates a new task.
	// The owner of the task is inferred from the authorizer associated with ctx.
	CreateTask(ctx context.Context, t TaskCreate) (*Task, error)

	// UpdateTask updates a single task with changeset.
	UpdateTask(ctx context.Context, id platform.ID, upd TaskUpdate) (*Task, error)

	// DeleteTask removes a task by ID and purges all associated data and scheduled runs.
	DeleteTask(ctx context.Context, id platform.ID) error

	// FindLogs returns logs for a run.
	FindLogs(ctx context.Context, filter LogFilter) ([]*Log, int, error)

	// FindRuns returns a list of runs that match a filter and the total count of returned runs.
	FindRuns(ctx context.Context, filter RunFilter) ([]*Run, int, error)

	// FindRunByID returns a single run.
	FindRunByID(ctx context.Context, taskID, runID platform.ID) (*Run, error)

	// CancelRun cancels a currently running run.
	CancelRun(ctx context.Context, taskID, runID platform.ID) error

	// RetryRun creates and returns a new run (which is a retry of another run).
	RetryRun(ctx context.Context, taskID, runID platform.ID) (*Run, error)

	// ForceRun forces a run to occur with unix timestamp scheduledFor, to be executed as soon as possible.
	// The value of scheduledFor may or may not align with the task's schedule.
	ForceRun(ctx context.Context, taskID platform.ID, scheduledFor int64) (*Run, error)
}

// TaskCreate is the set of values to create a task.
type TaskCreate struct {
	Type           string                 `json:"type,omitempty"`
	Flux           string                 `json:"flux"`
	Description    string                 `json:"description,omitempty"`
	Status         string                 `json:"status,omitempty"`
	OrganizationID platform.ID            `json:"orgID,omitempty"`
	Organization   string                 `json:"org,omitempty"`
	OwnerID        platform.ID            `json:"-"`
	Metadata       map[string]interface{} `json:"-"` // not to be set through a web request but rather used by a http service using tasks backend.
}

func (t TaskCreate) Validate() error {
	//switch {
	//case t.Flux == "":
	//	return errors.New("missing flux")
	//case !t.OrganizationID.Valid() && t.Organization == "":
	//	return errors.New("missing orgID and org")
	//case t.Status != "" && t.Status != TaskStatusActive && t.Status != TaskStatusInactive:
	//	return fmt.Errorf("invalid task status: %q", t.Status)
	//}
	return nil
}

// TaskUpdate represents updates to a task. Options updates override any options set in the Flux field.
type TaskUpdate struct {
	Flux        *string `json:"flux,omitempty"`
	Status      *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`

	// LatestCompleted us to set latest completed on startup to skip task catchup
	LatestCompleted *time.Time             `json:"-"`
	LatestScheduled *time.Time             `json:"-"`
	LatestSuccess   *time.Time             `json:"-"`
	LatestFailure   *time.Time             `json:"-"`
	LastRunStatus   *string                `json:"-"`
	LastRunError    *string                `json:"-"`
	Metadata        map[string]interface{} `json:"-"` // not to be set through a web request but rather used by a http service using tasks backend.

	// Options gets unmarshalled from json as if it was flat, with the same level as Flux and Status.
	Options options.Options // when we unmarshal this gets unmarshalled from flat key-values
}

func (t *TaskUpdate) UnmarshalJSON(data []byte) error {
	// this is a type so we can marshal string into durations nicely
	jo := struct {
		Flux        *string `json:"flux,omitempty"`
		Status      *string `json:"status,omitempty"`
		Name        string  `json:"name,omitempty"`
		Description *string `json:"description,omitempty"`

		// Cron is a cron style time schedule that can be used in place of Every.
		Cron string `json:"cron,omitempty"`

		// Every represents a fixed period to repeat execution.
		// It gets marshalled from a string duration, i.e.: "10s" is 10 seconds
		Every options.Duration `json:"every,omitempty"`

		// Offset represents a delay before execution.
		// It gets marshalled from a string duration, i.e.: "10s" is 10 seconds
		Offset *options.Duration `json:"offset,omitempty"`

		Concurrency *int64 `json:"concurrency,omitempty"`

		Retry *int64 `json:"retry,omitempty"`
	}{}

	if err := json.Unmarshal(data, &jo); err != nil {
		return err
	}
	t.Options.Name = jo.Name
	t.Description = jo.Description
	t.Options.Cron = jo.Cron
	t.Options.Every = jo.Every
	if jo.Offset != nil {
		offset := *jo.Offset
		t.Options.Offset = &offset
	}
	t.Options.Concurrency = jo.Concurrency
	t.Options.Retry = jo.Retry
	t.Flux = jo.Flux
	t.Status = jo.Status
	return nil
}

func (t *TaskUpdate) MarshalJSON() ([]byte, error) {
	jo := struct {
		Flux        *string `json:"flux,omitempty"`
		Status      *string `json:"status,omitempty"`
		Name        string  `json:"name,omitempty"`
		Description *string `json:"description,omitempty"`

		// Cron is a cron style time schedule that can be used in place of Every.
		Cron string `json:"cron,omitempty"`

		// Every represents a fixed period to repeat execution.
		Every options.Duration `json:"every,omitempty"`

		// Offset represents a delay before execution.
		Offset *options.Duration `json:"offset,omitempty"`

		Concurrency *int64 `json:"concurrency,omitempty"`

		Retry *int64 `json:"retry,omitempty"`
	}{}
	jo.Name = t.Options.Name
	jo.Cron = t.Options.Cron
	jo.Every = t.Options.Every
	jo.Description = t.Description
	if t.Options.Offset != nil {
		offset := *t.Options.Offset
		jo.Offset = &offset
	}
	jo.Concurrency = t.Options.Concurrency
	jo.Retry = t.Options.Retry
	jo.Flux = t.Flux
	jo.Status = t.Status
	return json.Marshal(jo)
}

func (t *TaskUpdate) Validate() error {
	switch {
	case !t.Options.Every.IsZero() && t.Options.Cron != "":
		return errors.New("cannot specify both every and cron")
	case !t.Options.Every.IsZero():
		if _, err := options.ParseSignedDuration(t.Options.Every.String()); err != nil {
			return fmt.Errorf("every: %s is invalid", err)
		}
	case t.Options.Offset != nil && !t.Options.Offset.IsZero():
		if _, err := time.ParseDuration(t.Options.Offset.String()); err != nil {
			return fmt.Errorf("offset: %s, %s is invalid, the largest unit supported is h", t.Options.Offset.String(), err)
		}
	case t.Flux == nil && t.Status == nil && t.Options.IsZero():
		return errors.New("cannot update task without content")
	case t.Status != nil && *t.Status != TaskStatusActive && *t.Status != TaskStatusInactive:
		return fmt.Errorf("invalid task status: %q", *t.Status)
	}
	return nil
}

// TaskFilter represents a set of filters that restrict the returned results
type TaskFilter struct {
	Type           *string
	Name           *string
	After          *platform.ID
	OrganizationID *platform.ID
	Organization   string
	User           *platform.ID
	Limit          int
	Status         *string
}

// QueryParams Converts TaskFilter fields to url query params.
func (f TaskFilter) QueryParams() map[string][]string {
	qp := map[string][]string{}
	if f.After != nil {
		qp["after"] = []string{f.After.String()}
	}

	if f.OrganizationID != nil {
		qp["orgID"] = []string{f.OrganizationID.String()}
	}

	if f.Organization != "" {
		qp["org"] = []string{f.Organization}
	}

	if f.User != nil {
		qp["user"] = []string{f.User.String()}
	}

	if f.Limit > 0 {
		qp["limit"] = []string{strconv.Itoa(f.Limit)}
	}

	return qp
}

// RunFilter represents a set of filters that restrict the returned results
type RunFilter struct {
	// Task ID is required for listing runs.
	Task platform.ID

	After      *platform.ID
	Limit      int
	AfterTime  string
	BeforeTime string
}

// LogFilter represents a set of filters that restrict the returned log results.
type LogFilter struct {
	// Task ID is required.
	Task platform.ID

	// The optional Run ID limits logs to a single run.
	Run *platform.ID
}

type TaskStatus string

const (
	TaskActive   TaskStatus = "active"
	TaskInactive TaskStatus = "inactive"

	DefaultTaskStatus TaskStatus = TaskActive
)

type RunStatus int

const (
	RunStarted RunStatus = iota
	RunSuccess
	RunFail
	RunCanceled
	RunScheduled
)

func (r RunStatus) String() string {
	switch r {
	case RunStarted:
		return "started"
	case RunSuccess:
		return "success"
	case RunFail:
		return "failed"
	case RunCanceled:
		return "canceled"
	case RunScheduled:
		return "scheduled"
	}
	panic(fmt.Sprintf("unknown RunStatus: %d", r))
}

// RequestStillQueuedError is returned when attempting to retry a run which has not yet completed.
type RequestStillQueuedError struct {
	// Unix timestamps matching existing request's start and end.
	Start, End int64
}

const fmtRequestStillQueued = "previous retry for start=%s end=%s has not yet finished"

func (e RequestStillQueuedError) Error() string {
	return fmt.Sprintf(fmtRequestStillQueued,
		time.Unix(e.Start, 0).UTC().Format(time.RFC3339),
		time.Unix(e.End, 0).UTC().Format(time.RFC3339),
	)
}

// ParseRequestStillQueuedError attempts to parse a RequestStillQueuedError from msg.
// If msg is formatted correctly, the resultant error is returned; otherwise it returns nil.
func ParseRequestStillQueuedError(msg string) *RequestStillQueuedError {
	var s, e string
	n, err := fmt.Sscanf(msg, fmtRequestStillQueued, &s, &e)
	if err != nil || n != 2 {
		return nil
	}

	start, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil
	}

	end, err := time.Parse(time.RFC3339, e)
	if err != nil {
		return nil
	}

	return &RequestStillQueuedError{Start: start.Unix(), End: end.Unix()}
}
