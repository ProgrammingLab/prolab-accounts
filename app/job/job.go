package job

import (
	"context"
	"time"

	"google.golang.org/grpc/grpclog"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
)

var (
	started = false
	stop    = make(chan struct{})
	jobs    = []Job{
		feedJob,
	}
)

// Job represents job for worker
type Job func(ctx context.Context, store di.StoreComponent) error

// Start starts the worker
func Start(store di.StoreComponent, cfg *config.Config) {
	if started {
		return
	}
	started = true

	go func() {
		run(store, time.Duration(cfg.JobIntervalSec)*time.Second)
	}()
}

// Close stops the worker
func Close() {
	grpclog.Infoln("worker is stopping(^C to force to stop)")
	stop <- struct{}{}
}

func run(store di.StoreComponent, interval time.Duration) {
	defer func() {
		if err := util.ErrorFromRecover(recover()); err != nil {
			grpclog.Errorf("job panic: %+v", err)
			grpclog.Infoln("worker is restarting...")
			run(store, interval)
		}
	}()

	grpclog.Infof("worker started: interval %v", interval)

	for {
		select {
		case <-time.After(interval):
			for _, j := range jobs {
				err := j(context.Background(), store)
				if err != nil {
					grpclog.Errorf("job error: %+v", err)
				}
			}
		case <-stop:
			grpclog.Infoln("worker stopped")
			return
		}
	}
}
