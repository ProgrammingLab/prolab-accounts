package server

import (
	"context"
	"database/sql"
	"fmt"
	"image"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	validator "github.com/mwitkow/go-proto-validators"
	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// NewUserServiceServer creates a new UserServiceServer instance.
func NewUserServiceServer(store di.StoreComponent, cfg *config.Config) interface {
	api_pb.UserServiceServer
	grapiserver.Server
} {
	return &userServiceServerImpl{
		StoreComponent: store,
		cfg:            cfg,
	}
}

type userServiceServerImpl struct {
	di.StoreComponent
	cfg *config.Config
}

const (
	// MaxIconSize represents max of icon size
	MaxIconSize = 1024 * 1024 // 1MiB
)

var (
	// ErrPageSizeOutOfRange will be returned when page size is out of range
	ErrPageSizeOutOfRange = status.Error(codes.OutOfRange, "page size must be 1 <= size <= 100")
	// ErrIconSizeTooLarge will be returned when icon is too large
	ErrIconSizeTooLarge = validator.FieldError("Image", fmt.Errorf("image must be smaller than 1MiB"))
	// ErrInvalidImageFormat will be returned when image format is invalid
	ErrInvalidImageFormat = status.Error(codes.InvalidArgument, "invalid iamge format")
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
		Users:         usersToResponse(u, false, s.cfg),
		NextPageToken: uint32(next),
	}
	return resp, nil
}

func (s *userServiceServerImpl) GetUser(ctx context.Context, req *api_pb.GetUserRequest) (*api_pb.User, error) {
	us := s.UserStore(ctx)
	u, err := us.GetPublicUserByName(req.GetUserName())
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return userToResponse(u, false, s.cfg), nil
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

	u, err := s.UserStore(ctx).GetUserWithPrivate(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.ErrUnauthenticated
		}
		return nil, err
	}

	return userToResponse(u, true, s.cfg), nil
}

func (s *userServiceServerImpl) UpdateUserProfile(ctx context.Context, req *api_pb.UpdateUserProfileRequest) (*api_pb.User, error) {
	id, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	us := s.UserStore(ctx)
	u, err := us.UpdateFullName(id, req.GetFullName())
	if err != nil {
		return nil, err
	}

	ps := s.ProfileStore(ctx)
	p := &record.Profile{
		ID:                u.ProfileID.Int64,
		Description:       req.GetDescription(),
		Grade:             int(req.GetGrade()),
		Left:              req.GetLeft(),
		TwitterScreenName: null.StringFrom(req.GetTwitterScreenName()),
		GithubUserName:    null.StringFrom(req.GetGithubUserName()),
		ProfileScope:      null.IntFrom(int(req.GetProfileScope())),
	}
	if id := req.GetRoleId(); id == 0 {
		p.RoleID = null.NewInt64(0, false)
	} else {
		p.RoleID = null.Int64From(int64(id))
	}
	if id := req.GetDepartmentId(); id == 0 {
		p.DepartmentID = null.NewInt64(0, false)
	} else {
		p.DepartmentID = null.Int64From(int64(id))
	}

	err = ps.CreateOrUpdateProfile(model.UserID(u.ID), p)
	if err != nil {
		return nil, err
	}

	u, err = us.GetUserWithPrivate(id)
	if err != nil {
		return nil, err
	}

	return userToResponse(u, true, s.cfg), nil
}

func (s *userServiceServerImpl) UpdateUserIcon(ctx context.Context, req *api_pb.UpdateUserIconRequest) (*api_pb.User, error) {
	id, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	icon := req.GetImage()
	if MaxIconSize < len(icon) {
		return nil, ErrIconSizeTooLarge
	}

	us := s.UserStore(ctx)
	u, err := us.UpdateIcon(id, icon)
	if err != nil {
		if err := errors.Cause(err); err == image.ErrFormat {
			return nil, ErrInvalidImageFormat
		}
		return nil, err
	}

	return userToResponse(u, true, s.cfg), nil
}

func (s *userServiceServerImpl) UpdatePassword(ctx context.Context, req *api_pb.UpdatePasswordRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func usersToResponse(users []*record.User, includePrivate bool, cfg *config.Config) []*api_pb.User {
	res := make([]*api_pb.User, 0, len(users))
	for _, u := range users {
		res = append(res, userToResponse(u, includePrivate, cfg))
	}
	return res
}

func userToResponse(user *record.User, includePrivate bool, cfg *config.Config) *api_pb.User {
	if user == nil {
		return nil
	}

	u := &api_pb.User{
		UserId: uint32(user.ID),
		Name:   user.Name,
	}
	if includePrivate {
		u.Email = user.Email
		u.FullName = user.FullName
	}
	if user.AvatarFilename.Valid {
		u.IconUrl = cfg.MinioPublicURL + "/" + cfg.MinioBucketName + "/" + user.AvatarFilename.String
	}
	if r := user.R; r != nil && r.Profile != nil {
		p := r.Profile
		u.Description = p.Description
		u.Grade = int32(p.Grade)
		u.Left = p.Left
		u.TwitterScreenName = p.TwitterScreenName.String
		u.GithubUserName = p.GithubUserName.String
		u.ProfileScope = profileScopeToResponse(model.ProfileScope(p.ProfileScope.Int))

		if r := p.R; p.R != nil {
			if role := r.Role; role != nil {
				u.Role = roleToResponse(role)
			}
			if dep := r.Department; dep != nil {
				u.Department = departmentToResponse(dep)
			}
		}
	}

	return u
}

func profileScopeToResponse(scope model.ProfileScope) api_pb.ProfileScope {
	switch scope {
	case model.MembersOnly:
		return api_pb.ProfileScope_MEMBERS_ONLY
	case model.Public:
		return api_pb.ProfileScope_PUBLIC
	default:
		//unknow
		return -1
	}
}
