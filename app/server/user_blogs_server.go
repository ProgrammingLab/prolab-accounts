package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// UserBlogServiceServer is a composite interface of api_pb.UserBlogServiceServer and grapiserver.Server.
type UserBlogServiceServer interface {
	api_pb.UserBlogServiceServer
	grapiserver.Server
}

// NewUserBlogServiceServer creates a new UserBlogServiceServer instance.
func NewUserBlogServiceServer() UserBlogServiceServer {
	return &userBlogServiceServerImpl{}
}

type userBlogServiceServerImpl struct {
}

func (s *userBlogServiceServerImpl) ListUserBlogs(ctx context.Context, req *api_pb.ListUserBlogsRequest) (*api_pb.ListUserBlogsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userBlogServiceServerImpl) UpdateUserBlogs(ctx context.Context, req *api_pb.UpdateUserBlogsRequest) (*api_pb.ListUserBlogsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
