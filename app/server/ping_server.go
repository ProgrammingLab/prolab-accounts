package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

var (
	// ErrWorkerStopped retuned when the worker has stopped
	ErrWorkerStopped = status.Error(codes.Internal, "worker has stopped")
)

func (s *pingServiceServerImpl) Ping(ctx context.Context, req *api_pb.PingRequest) (*api_pb.Pong, error) {
	hs := s.HeartbeatStore(ctx)
	err := hs.GetHeartbeat()
	if err != nil {
		return nil, ErrWorkerStopped
	}
	return &api_pb.Pong{}, nil
}
