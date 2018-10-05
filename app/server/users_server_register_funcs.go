package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *userServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterUserServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *userServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterUserServiceHandler(ctx, mux, conn)
}
