package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

// UserBlogServiceServer is a composite interface of api_pb.UserBlogServiceServer and grapiserver.Server.
type UserBlogServiceServer interface {
	api_pb.UserBlogServiceServer
	grapiserver.Server
}

// NewUserBlogServiceServer creates a new UserBlogServiceServer instance.
func NewUserBlogServiceServer(store di.StoreComponent) UserBlogServiceServer {
	return &userBlogServiceServerImpl{
		StoreComponent: store,
	}
}

type userBlogServiceServerImpl struct {
	di.StoreComponent
}

func (s *userBlogServiceServerImpl) ListUserBlogs(ctx context.Context, req *api_pb.ListUserBlogsRequest) (*api_pb.ListUserBlogsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userBlogServiceServerImpl) GetUserBlog(ctx context.Context, req *api_pb.GetUserBlogRequest) (*api_pb.Blog, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userBlogServiceServerImpl) CreateUserBlog(ctx context.Context, req *api_pb.CreateUserBlogRequest) (*api_pb.Blog, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userBlogServiceServerImpl) UpdateUserBlog(ctx context.Context, req *api_pb.UpdateUserBlogRequest) (*api_pb.Blog, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userBlogServiceServerImpl) DeleteUserBlog(ctx context.Context, req *api_pb.DeleteUserBlogRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
