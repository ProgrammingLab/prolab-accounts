package di

import (
	"context"
	"net/url"

	"github.com/ory/hydra/sdk/go/hydra/client"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/infra/email"
	"github.com/ProgrammingLab/prolab-accounts/static"
)

// ClientComponent is an interface of api clients
type ClientComponent interface {
	HydraClient(ctx context.Context) *client.OryHydra
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

func newHydraClient(cfg *config.Config) (*client.OryHydra, error) {
	adminURL, err := url.Parse(cfg.HydraAdminURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	cliCfg := &client.TransportConfig{
		Schemes:  []string{adminURL.Scheme},
		Host:     adminURL.Host,
		BasePath: adminURL.Path,
	}
	cli := client.NewHTTPClientWithConfig(nil, cliCfg)
	return cli, nil
}

type clientComponentImpl struct {
	cfg      *config.Config
	hydraCli *client.OryHydra
	emails   *static.EmailAsset
}

func (c *clientComponentImpl) HydraClient(ctx context.Context) *client.OryHydra {
	return c.hydraCli
}

func (c *clientComponentImpl) EmailSender(ctx context.Context) email.Sender {
	return email.NewSender(ctx, c.cfg, c.emails)
}
