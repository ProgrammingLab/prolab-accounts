package userblogstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type userBlogStoreImpl struct {
	ctx context.Context
	db  *sqlutil.DB
}

// NewUserBlogStore returns new user blog store
func NewUserBlogStore(ctx context.Context, db *sqlutil.DB) store.UserBlogStore {
	return &userBlogStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *userBlogStoreImpl) ListUserBlogs() ([]*record.Blog, error) {
	b, err := record.Blogs().All(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return b, nil
}

func (s *userBlogStoreImpl) GetUserBlog(blogID int64) (*record.Blog, error) {
	b, err := record.FindBlog(s.ctx, s.db, blogID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return b, nil
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
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		exists, err := record.FindBlog(s.ctx, tx, blog.ID)
		if err != nil {
			return errors.WithStack(err)
		}
		if exists.UserID != blog.UserID {
			return sql.ErrNoRows
		}

		_, err = blog.Update(s.ctx, tx, boil.Infer())
		return errors.WithStack(err)
	})
	return err
}

func (s *userBlogStoreImpl) DeleteUserBlog(blogID int64) error {
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, err := record.Entries(record.EntryWhere.BlogID.EQ(blogID)).DeleteAll(s.ctx, tx)
		if err != nil {
			_ = tx.Rollback()
			return errors.WithStack(err)
		}

		_, err = record.Blogs(record.BlogWhere.ID.EQ(blogID)).DeleteAll(s.ctx, tx)
		return errors.WithStack(err)
	})
	return err
}
