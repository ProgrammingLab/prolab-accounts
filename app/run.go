package app

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/rs/cors"
	"github.com/srvc/appctx"
	"github.com/volatiletech/sqlboiler/boil"
	"google.golang.org/grpc/grpclog"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/job"
	"github.com/ProgrammingLab/prolab-accounts/app/server"
)

// Run starts the grapiserver.
func Run() error {
	ctx := appctx.Global()

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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*.prolab.club", "https://prolab.club", "http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept"},
		AllowCredentials: true,
		Debug:            cfg.DebugLog,
	})

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGatewayServerMiddlewares(
			c.Handler,
		),
		grapiserver.WithGatewayMuxOptions(
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		),
		grapiserver.WithGrpcServerUnaryInterceptors(
			interceptor.RecoverUnaryServerInterceptor(),
			interceptor.ErrorUnaryServerInterceptor(),
			interceptor.ValidationUnaryServerInterceptor(),
			authorizator.UnaryServerInterceptor(),
		),
		grapiserver.WithServers(
			server.NewSessionServiceServer(store),
			server.NewUserServiceServer(store, cfg),
			server.NewOAuthServiceServer(cli, store),
			server.NewUserBlogServiceServer(store),
			server.NewEntryServiceServer(store, cfg),
			server.NewPingServiceServer(store),
			server.NewRoleServiceServer(store),
			server.NewDepartmentServiceServer(store),
			server.NewInvitationServiceServer(store, cli),
			server.NewContributionCollectionServiceServer(store, cfg),
			server.NewEmailConfirmationServiceServer(store, cli, cfg),
			server.NewPasswordResetServiceServer(store, cli, cfg),
			server.NewAchievementServiceServer(store, cli, cfg),
		),
	)

	job.Start(store, cfg)
	defer job.Close()

	go func() {
		err := store.ImageStore(context.TODO()).MigrateImages()
		if err != nil {
			grpclog.Errorf("%+v", err)
		}
	}()

	return s.Serve(ctx)
}
