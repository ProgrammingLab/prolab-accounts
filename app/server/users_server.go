package server

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/dao"
)

// NewUserServiceServer creates a new UserServiceServer instance.
func NewUserServiceServer(store di.StoreComponent) interface {
	api_pb.UserServiceServer
	grapiserver.Server
} {
	return &userServiceServerImpl{
		StoreComponent: store,
	}
}

type userServiceServerImpl struct {
	di.StoreComponent
}

func (s *userServiceServerImpl) ListUsers(ctx context.Context, req *api_pb.ListUsersRequest) (*api_pb.ListUsersResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) GetUser(ctx context.Context, req *api_pb.GetUserRequest) (*api_pb.User, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) CreateUser(ctx context.Context, req *api_pb.CreateUserRequest) (*api_pb.User, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) GetCurrentUser(ctx context.Context, req *api_pb.GetCurrentUserRequest) (*api_pb.User, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	u, err := s.UserStore(ctx).GetUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.ErrUnauthenticated
		}
		return nil, err
	}

	return userToResponse(u, true), nil
}

func (s *userServiceServerImpl) UpdateUser(ctx context.Context, req *api_pb.UpdateUserRequest) (*api_pb.User, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) UpdatePassword(ctx context.Context, req *api_pb.UpdatePasswordRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func userToResponse(user *dao.User, includeEmail bool) *api_pb.User {
	var email string
	if includeEmail {
		email = user.Email
	}

	return &api_pb.User{
		UserId:    uint32(user.ID),
		Name:      user.Name,
		Email:     email,
		FullName:  user.FullName,
		AvatarUrl: "not implemented",
	}
}
