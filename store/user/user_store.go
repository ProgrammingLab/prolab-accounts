package userstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ProgrammingLab/prolab-accounts/infra/model"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/store"
)

type userStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewUserStore returns new user store
func NewUserStore(ctx context.Context, db *sql.DB) store.UserStore {
	return &userStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *userStoreImpl) CreateUser(user *record.User) error {
	err := user.Insert(s.ctx, s.db, boil.Infer())
	return errors.WithStack(err)
}

func (s *userStoreImpl) GetUser(userID model.UserID) (*record.User, error) {
	u, err := record.Users(qm.Where("id = ?", userID)).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "")
	}

	return u, nil
}
