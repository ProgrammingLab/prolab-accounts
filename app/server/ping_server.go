package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

// PingServiceServer is a composite interface of api_pb.PingServiceServer and grapiserver.Server.
type PingServiceServer interface {
	api_pb.PingServiceServer
	grapiserver.Server
}

// NewPingServiceServer creates a new PingServiceServer instance.
func NewPingServiceServer(store di.StoreComponent) PingServiceServer {
	return &pingServiceServerImpl{
		StoreComponent: store,
	}
}

type pingServiceServerImpl struct {
	di.StoreComponent
}

func (s *pingServiceServerImpl) Ping(ctx context.Context, req *api_pb.PingRequest) (*api_pb.Pong, error) {
	return &api_pb.Pong{}, nil
}
