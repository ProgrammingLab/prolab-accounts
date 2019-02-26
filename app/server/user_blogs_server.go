package server

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
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

func (s *userBlogServiceServerImpl) FindFeedURL(ctx context.Context, req *api_pb.FindFeedURLRequest) (*api_pb.Blog, error) {
	_, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	fs := s.FeedStore(ctx)
	u, err := fs.GetFeedURL(req.GetUrl())
	if err != nil {
		return nil, ErrFeedURLDetectAutomatically
	}
	return &api_pb.Blog{
		Url:     req.GetUrl(),
		FeedUrl: u,
	}, nil
}

func (s *userBlogServiceServerImpl) CreateUserBlog(ctx context.Context, req *api_pb.CreateUserBlogRequest) (*api_pb.Blog, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	blog := req.GetBlog()
	b := &record.Blog{
		URL:     blog.GetUrl(),
		FeedURL: blog.GetFeedUrl(),
		UserID:  int64(userID),
	}

	fs := s.FeedStore(ctx)
	_, err := fs.GetFeed(b.FeedURL)
	if err != nil {
		return nil, ErrInvalidFeedURL
	}

	bs := s.UserBlogStore(ctx)
	err = bs.CreateUserBlog(b)
	if err != nil {
		return nil, err
	}

	return blogToResponse(b), nil
}

func (s *userBlogServiceServerImpl) UpdateUserBlog(ctx context.Context, req *api_pb.UpdateUserBlogRequest) (*api_pb.Blog, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	blog := req.GetBlog()
	b := &record.Blog{
		ID:      int64(blog.GetBlogId()),
		URL:     blog.GetUrl(),
		FeedURL: blog.GetFeedUrl(),
		UserID:  int64(userID),
	}

	fs := s.FeedStore(ctx)
	_, err := fs.GetFeed(b.FeedURL)
	if err != nil {
		return nil, ErrInvalidFeedURL
	}

	if err := s.canWrite(ctx, userID, b.ID); err != nil {
		return nil, err
	}

	bs := s.UserBlogStore(ctx)
	err = bs.UpdateUserBlog(b)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return blogToResponse(b), nil
}

func (s *userBlogServiceServerImpl) DeleteUserBlog(ctx context.Context, req *api_pb.DeleteUserBlogRequest) (*empty.Empty, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	blogID := int64(req.GetBlogId())
	bs := s.UserBlogStore(ctx)

	if err := s.canWrite(ctx, userID, blogID); err != nil {
		return nil, err
	}

	err := bs.DeleteUserBlog(blogID)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *userBlogServiceServerImpl) canWrite(ctx context.Context, userID model.UserID, blogID int64) error {
	bs := s.UserBlogStore(ctx)
	b, err := bs.GetUserBlog(int64(blogID))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return util.ErrNotFound
		}
		return err
	}
	if b.UserID != int64(userID) {
		return util.ErrNotFound
	}
	return nil
}

type blogRequest interface {
	GetBlog() *api_pb.Blog
	GetAutoDetectFeed() bool
}

func blogToResponse(blog *record.Blog) *api_pb.Blog {
	if blog == nil {
		return nil
	}

	return &api_pb.Blog{
		BlogId:  uint32(blog.ID),
		Url:     blog.URL,
		FeedUrl: blog.FeedURL,
	}
}
