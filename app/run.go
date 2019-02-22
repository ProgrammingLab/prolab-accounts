package app

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/volatiletech/sqlboiler/boil"
	"google.golang.org/grpc/grpclog"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/server"
)

// Run starts the grapiserver.
func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		grpclog.Errorf("%+v", err)
		return err
	}

	store, err := di.NewStoreComponent(cfg)
	if err != nil {
		grpclog.Errorf("%+v", err)
		return err
	}

	cli, err := di.NewClientComponent(cfg)
	if err != nil {
		grpclog.Errorf("%+v", err)
		return err
	}

	boil.DebugMode = cfg.DebugLog

	authorizator := interceptor.NewAuthorizator(store)

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGatewayServerMiddlewares(
			interceptor.CORSMiddleware,
		),
		grapiserver.WithGatewayMuxOptions(
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		),
		grapiserver.WithGrpcServerUnaryInterceptors(
			interceptor.ErrorUnaryServerInterceptor(),
			interceptor.ValidationUnaryServerInterceptor(),
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
