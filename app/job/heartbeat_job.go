package job

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

func heartbeatJob(ctx context.Context, store di.StoreComponent, cfg *config.Config) error {
	s := store.HeartbeatStore(ctx)
	return errors.WithStack(s.Beat())
}
