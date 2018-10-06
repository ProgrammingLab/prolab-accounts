package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *userProfileServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterUserProfileServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *userProfileServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterUserProfileServiceHandler(ctx, mux, conn)
}
