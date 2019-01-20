package app

import (
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/labstack/gommon/log"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/server"
)

// Run starts the grapiserver.
func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
		return err
	}

	store, err := di.NewStoreComponent(cfg)
	if err != nil {
		log.Error(err)
		return err
	}

	cli, err := di.NewClientComponent(cfg)
	if err != nil {
		log.Error(err)
		return err
	}

	boil.DebugMode = cfg.DebugLog

	authorizator := interceptor.NewAuthorizator(store)

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGrpcServerUnaryInterceptors(
			authorizator.UnaryServerInterceptor(),
		),
		grapiserver.WithServers(
			server.NewSessionServiceServer(store),
			server.NewUserServiceServer(store),
			server.NewOAuthServiceServer(cli, store),
		),
	)
	return s.Serve()
}
