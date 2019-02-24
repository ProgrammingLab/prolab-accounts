package userblogstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
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

func (s *userBlogStoreImpl) CreateUserBlog(blog *record.Blog) error {
	blog.ID = 0
	err := blog.Insert(s.ctx, s.db, boil.Infer())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *userBlogStoreImpl) UpdateUserBlog(blog *record.Blog) error {
	panic("not implemented")
}

func (s *userBlogStoreImpl) DeleteUserBlog(blogID int64) error {
	panic("not implemented")
}
