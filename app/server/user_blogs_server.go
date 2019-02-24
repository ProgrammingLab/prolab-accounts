package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
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

var (
	// ErrFeedURLDetectAutomatically returns when feed url could not be found automatically
	ErrFeedURLDetectAutomatically = status.Error(codes.InvalidArgument, "feed url could not be found automatically")
	// ErrInvalidFeedURL returns when feed url is invalid
	ErrInvalidFeedURL = status.Error(codes.InvalidArgument, "feed url is invalid")
)

func (s *userBlogServiceServerImpl) CreateUserBlog(ctx context.Context, req *api_pb.CreateUserBlogRequest) (*api_pb.Blog, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	blog := req.GetBlog()
	var feedURL string
	if req.GetAutoDetectFeed() {
		fs := s.FeedStore(ctx)
		u, err := fs.GetFeedURL(blog.GetUrl())
		if err != nil {
			return nil, ErrFeedURLDetectAutomatically
		}
		feedURL = u
	} else {
		u := blog.GetFeedUrl()
		fs := s.FeedStore(ctx)
		err := fs.IsValidFeedURL(u)
		if err != nil {
			return nil, ErrInvalidFeedURL
		}
		feedURL = u
	}

	b := &record.Blog{
		URL:     blog.GetUrl(),
		FeedURL: feedURL,
		UserID:  int64(userID),
	}

	bs := s.UserBlogStore(ctx)
	err := bs.CreateUserBlog(b)
	if err != nil {
		return nil, err
	}

	return blogToResponse(b), nil
}

func (s *userBlogServiceServerImpl) UpdateUserBlog(ctx context.Context, req *api_pb.UpdateUserBlogRequest) (*api_pb.Blog, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userBlogServiceServerImpl) DeleteUserBlog(ctx context.Context, req *api_pb.DeleteUserBlogRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func blogToResponse(blog *record.Blog) *api_pb.Blog {
	return &api_pb.Blog{
		BlogId:  uint32(blog.ID),
		Url:     blog.URL,
		FeedUrl: blog.FeedURL,
	}
}
