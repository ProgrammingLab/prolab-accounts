package heartbeatstore

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
)

type heartbeatStoreImpl struct {
	ctx    context.Context
	client *redis.Client
	cfg    *config.Config
}

// NewHeartbeatStore returns new heartbeat store
func NewHeartbeatStore(ctx context.Context, cli *redis.Client, cfg *config.Config) store.HeartbeatStore {
	return &heartbeatStoreImpl{
		ctx:    ctx,
		client: cli,
		cfg:    cfg,
	}
}

const key = "heartbeat"

func (s *heartbeatStoreImpl) Beat() error {
	exp := time.Duration(s.cfg.JobIntervalSec) * time.Second * 2
	err := s.client.Set(key, "dokidoki", exp).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *heartbeatStoreImpl) GetHeartbeat() error {
	return errors.WithStack(s.client.Get(key).Err())
}
