package emailconfirmationstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type emailConfirmationStoreImpl struct {
	ctx context.Context
	db  *sqlutil.DB
}

// NewEmailConfirmationStore returns new email confirmation store
func NewEmailConfirmationStore(ctx context.Context, db *sqlutil.DB) store.EmailConfirmationStore {
	return &emailConfirmationStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

const (
	tokenLength = 32
)

func (s *emailConfirmationStoreImpl) CreateConfirmation(userID model.UserID, email string) (*record.EmailConfirmation, error) {
	t, err := model.GenerateSecureToken(tokenLength)
	if err != nil {
		return nil, err
	}

	c := &record.EmailConfirmation{
		Token:  t,
		UserID: int64(userID),
		Email:  email,
	}

	err = s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		ex, err := record.Users(record.UserWhere.Email.EQ(email)).Exists(ctx, tx)
		if err != nil {
			return errors.WithStack(err)
		}

		if ex {
			return store.ErrEmailAlreadyInUse
		}

		_, err = record.EmailConfirmations(record.EmailConfirmationWhere.Email.EQ(email)).DeleteAll(ctx, tx)
		if err != nil {
			return errors.WithStack(err)
		}

		err = c.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return errors.WithStack(err)
		}

		err = c.L.LoadUser(s.ctx, tx, true, c, nil)
		return errors.WithStack(err)
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *emailConfirmationStoreImpl) ConfirmEmail(token string) error {
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		c, err := record.EmailConfirmations(record.EmailConfirmationWhere.Token.EQ(token)).One(ctx, tx)
		if err != nil {
			return errors.WithStack(err)
		}

		u := &record.User{
			ID:    int64(c.UserID),
			Email: c.Email,
		}
		_, err = u.Update(ctx, tx, boil.Whitelist("email", "updated_at"))
		return errors.WithStack(err)
	})
	return err
}
