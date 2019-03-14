package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// PasswordResetServiceServer is a composite interface of api_pb.PasswordResetServiceServer and grapiserver.Server.
type PasswordResetServiceServer interface {
	api_pb.PasswordResetServiceServer
	grapiserver.Server
}

// NewPasswordResetServiceServer creates a new PasswordResetServiceServer instance.
func NewPasswordResetServiceServer() PasswordResetServiceServer {
	return &passwordResetServiceServerImpl{}
}

type passwordResetServiceServerImpl struct {
}

func (s *passwordResetServiceServerImpl) ListPasswordResets(ctx context.Context, req *api_pb.ListPasswordResetsRequest) (*api_pb.ListPasswordResetsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *passwordResetServiceServerImpl) GetPasswordReset(ctx context.Context, req *api_pb.GetPasswordResetRequest) (*api_pb.PasswordReset, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *passwordResetServiceServerImpl) CreatePasswordReset(ctx context.Context, req *api_pb.CreatePasswordResetRequest) (*api_pb.PasswordReset, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *passwordResetServiceServerImpl) UpdatePasswordReset(ctx context.Context, req *api_pb.UpdatePasswordResetRequest) (*api_pb.PasswordReset, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *passwordResetServiceServerImpl) DeletePasswordReset(ctx context.Context, req *api_pb.DeletePasswordResetRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
