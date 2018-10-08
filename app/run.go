package app

import (
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/labstack/gommon/log"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

// Run starts the grapiserver.
func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = di.NewStoreComponent(cfg)
	if err != nil {
		log.Error(err)
		return err
	}

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
		// TODO
		),
	)
	return s.Serve()
}
