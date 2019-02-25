package userblogstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
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

func (s *userBlogStoreImpl) ListUserBlogs() ([]*record.Blog, error) {
	b, err := record.Blogs().All(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return b, nil
}

func (s *userBlogStoreImpl) GetUserBlog(blogID int64) (*record.Blog, error) {
	b, err := record.FindBlog(s.ctx, s.db, int64(blogID))
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

func (s *userBlogStoreImpl) UpdateUserBlog(blog *record.Blog) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if e := util.ErrorFromRecover(recover()); e != nil {
			_ = tx.Rollback()
			err = e
		}
	}()

	exists, err := record.FindBlog(s.ctx, tx, blog.ID)
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}
	if exists.UserID != blog.UserID {
		_ = tx.Rollback()
		return sql.ErrNoRows
	}

	_, err = blog.Update(s.ctx, tx, boil.Infer())
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}

	return nil
}

func (s *userBlogStoreImpl) DeleteUserBlog(blogID int64) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if e := util.ErrorFromRecover(recover()); e != nil {
			_ = tx.Rollback()
			err = e
		}
	}()

	_, err = record.Entries(record.EntryWhere.BlogID.EQ(blogID)).DeleteAll(s.ctx, tx)
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}

	_, err = record.Blogs(record.BlogWhere.ID.EQ(blogID)).DeleteAll(s.ctx, tx)
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}

	return nil
}
