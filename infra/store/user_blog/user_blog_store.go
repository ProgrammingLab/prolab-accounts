package userblogstore

import (
	"context"
	"database/sql"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

type userBlogStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewUserBlogStore returns new user blog store
func NewUserBlogStore(ctx context.Context, db *sql.DB) store.UserBlogStore {
	return &userBlogStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *userBlogStoreImpl) CreateUserBlog(userID model.UserID, blog *record.Blog, detectFeed bool) error {
	panic("not implemented")
}

func (s *userBlogStoreImpl) UpdateUserBlog(userID model.UserID, blog *record.Blog, detectFeed bool) error {
	panic("not implemented")
}

func (s *userBlogStoreImpl) DeleteUserBlog(blogID int64) error {
	panic("not implemented")
}
