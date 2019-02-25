package job

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

func heartbeatJob(ctx context.Context, store di.StoreComponent, debug bool) error {
	s := store.HeartbeatStore(ctx)
	return errors.WithStack(s.Beat())
}
