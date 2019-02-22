package di

import (
	"context"

	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
)

// ClientComponent is an interface of api clients
type ClientComponent interface {
	HydraClient(ctx context.Context) *hydra.CodeGenSDK
}

// NewClientComponent returns new client component
func NewClientComponent(cfg *config.Config) (ClientComponent, error) {
	h, err := newHydraClient(cfg)
	if err != nil {
		return nil, err
	}

	return &clientComponentImpl{
		hydraCli: h,
	}, nil
}

func newHydraClient(cfg *config.Config) (*hydra.CodeGenSDK, error) {
	hc := &hydra.Configuration{
		AdminURL: cfg.HydraAdminURL,
	}
	cli, err := hydra.NewSDK(hc)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return cli, nil
}

type clientComponentImpl struct {
	hydraCli *hydra.CodeGenSDK
}

func (c *clientComponentImpl) HydraClient(ctx context.Context) *hydra.CodeGenSDK {
	return c.hydraCli
}
