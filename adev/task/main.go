package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb/v2/kit/signals"
	"github.com/uvite/gvmdesk/pkg/launcher"
	"time"
)

func vm() error {
	fmt.Println(2)
	l := launcher.NewLauncher()
	fmt.Println(3)
	ctx := context.Background()
	o := launcher.NewOpts()
	fmt.Printf("%+v", o)

	// Start the launcher and wait for it to exit on SIGINT or SIGTERM.
	if err := l.Run(signals.WithStandardSignals(ctx), o); err != nil {
		return err
	}
	<-l.Done()

	// Tear down the launcher, allowing it a few seconds to finish any
	// in-progress requests.
	shutdownCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return l.Shutdown(shutdownCtx)
}
func main() {
	fmt.Println(1)
	vm()
	//log := zap.NewNop()
	//
	//executor := task.NewExecutor(
	//	log.With(zap.String("service", "task-executor")),
	//	query.QueryServiceBridge{AsyncQueryService: m.queryController},
	//	ts.UserService,
	//
	//)
	//err = executor.LoadExistingScheduleRuns(ctx)
	//if err != nil {
	//	m.log.Fatal("could not load existing scheduled runs", zap.Error(err))
	//}
}
