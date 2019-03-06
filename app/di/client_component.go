package di

import (
	"context"

	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/infra/email"
	"github.com/ProgrammingLab/prolab-accounts/static"
)

// ClientComponent is an interface of api clients
type ClientComponent interface {
	HydraClient(ctx context.Context) *hydra.CodeGenSDK
	EmailSender(ctx context.Context) email.Sender
}

// NewClientComponent returns new client component
func NewClientComponent(cfg *config.Config) (ClientComponent, error) {
	h, err := newHydraClient(cfg)
	if err != nil {
		return nil, err
	}

	e, err := static.LoadEmailTemplates()
	if err != nil {
		return nil, err
	}

	return &clientComponentImpl{
		cfg:      cfg,
		hydraCli: h,
		emails:   e,
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
	cfg      *config.Config
	hydraCli *hydra.CodeGenSDK
	emails   *static.EmailAsset
}

func (c *clientComponentImpl) HydraClient(ctx context.Context) *hydra.CodeGenSDK {
	return c.hydraCli
}

func (c *clientComponentImpl) EmailSender(ctx context.Context) email.Sender {
	return email.NewSender(ctx, c.cfg, c.emails)
}
