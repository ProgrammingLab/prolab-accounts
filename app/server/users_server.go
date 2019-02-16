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
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
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

var (
	// ErrPageSizeOutOfRange will be returned when page size is out of range
	ErrPageSizeOutOfRange = status.Error(codes.OutOfRange, "page size must be 1 <= size <= 100")
)

func (s *userServiceServerImpl) ListPublicUsers(ctx context.Context, req *api_pb.ListUsersRequest) (*api_pb.ListUsersResponse, error) {
	size := req.GetPageSize()
	if size < 0 || 100 < size {
		return nil, ErrPageSizeOutOfRange
	}
	if size == 0 {
		size = 50
	}

	us := s.UserStore(ctx)
	u, next, err := us.ListPublicUsers(model.UserID(req.GetPageToken()), int(size))
	if err != nil {
		return nil, err
	}

	resp := &api_pb.ListUsersResponse{
		Users:         usersToResponse(u, false),
		NextPageToken: uint32(next),
	}
	return resp, nil
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

func usersToResponse(users []*record.User, includeEmail bool) []*api_pb.User {
	res := make([]*api_pb.User, 0, len(users))
	for _, u := range users {
		res = append(res, userToResponse(u, includeEmail))
	}
	return res
}

func userToResponse(user *record.User, includeEmail bool) *api_pb.User {
	var email string
	if includeEmail {
		email = user.Email
	}

	u := &api_pb.User{
		UserId:    uint32(user.ID),
		Name:      user.Name,
		Email:     email,
		FullName:  user.FullName,
		AvatarUrl: "not implemented",
	}
	if r := user.R; r != nil && r.Profile != nil {
		p := r.Profile
		u.Description = p.Description
		u.Grade = int32(p.Grade)
		u.Left = p.Left
		// TODO
		u.Department = "not implemented"
		// TODO
		u.ShortDepartment = "not implemented"
		if p.R != nil && p.R.Role != nil {
			u.Role = p.R.Role.Name.String
		}
		u.TwitterScreenName = p.TwitterScreenName.String
		u.GithubUserName = p.GithubUserName.String
	}

	return u
}
