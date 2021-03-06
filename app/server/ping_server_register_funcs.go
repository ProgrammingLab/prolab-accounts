// Code generated by github.com/izumin5210/grapi. DO NOT EDIT.

package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *pingServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterPingServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *pingServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterPingServiceHandler(ctx, mux, conn)
}
