package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// NewSessionServiceServer creates a new SessionServiceServer instance.
func NewSessionServiceServer() interface {
	api_pb.SessionServiceServer
	grapiserver.Server
} {
	return &sessionServiceServerImpl{}
}

type sessionServiceServerImpl struct {
}

func (s *sessionServiceServerImpl) ListSessions(ctx context.Context, req *api_pb.ListSessionsRequest) (*api_pb.ListSessionsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *sessionServiceServerImpl) GetSession(ctx context.Context, req *api_pb.GetSessionRequest) (*api_pb.Session, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *sessionServiceServerImpl) CreateSession(ctx context.Context, req *api_pb.CreateSessionRequest) (*api_pb.Session, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *sessionServiceServerImpl) UpdateSession(ctx context.Context, req *api_pb.UpdateSessionRequest) (*api_pb.Session, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *sessionServiceServerImpl) DeleteSession(ctx context.Context, req *api_pb.DeleteSessionRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
