package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

// PasswordResetServiceServer is a composite interface of api_pb.PasswordResetServiceServer and grapiserver.Server.
type PasswordResetServiceServer interface {
	api_pb.PasswordResetServiceServer
	grapiserver.Server
}

// NewPasswordResetServiceServer creates a new PasswordResetServiceServer instance.
func NewPasswordResetServiceServer(store di.StoreComponent, cfg *config.Config) PasswordResetServiceServer {
	return &passwordResetServiceServerImpl{
		StoreComponent: store,
		cfg:            cfg,
	}
}

type passwordResetServiceServerImpl struct {
	di.StoreComponent
	cfg *config.Config
}

func (s *passwordResetServiceServerImpl) GetPasswordReset(ctx context.Context, req *api_pb.GetPasswordResetRequest) (*api_pb.PasswordReset, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *passwordResetServiceServerImpl) CreatePasswordReset(ctx context.Context, req *api_pb.CreatePasswordResetRequest) (*api_pb.PasswordReset, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *passwordResetServiceServerImpl) UpdatePassword(ctx context.Context, req *api_pb.UpdatePasswordRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
