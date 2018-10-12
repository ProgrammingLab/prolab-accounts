package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// NewUserProfileServiceServer creates a new UserProfileServiceServer instance.
func NewUserProfileServiceServer() interface {
	api_pb.UserProfileServiceServer
	grapiserver.Server
} {
	return &userProfileServiceServerImpl{}
}

type userProfileServiceServerImpl struct {
}

func (s *userProfileServiceServerImpl) GetUserProfile(ctx context.Context, req *api_pb.GetUserProfileRequest) (*api_pb.UserProfile, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userProfileServiceServerImpl) UpdateUserProfile(ctx context.Context, req *api_pb.UpdateUserProfileRequest) (*api_pb.UserProfile, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
