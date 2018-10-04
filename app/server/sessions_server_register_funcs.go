package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *sessionServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterSessionServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *sessionServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterSessionServiceHandler(ctx, mux, conn)
}
