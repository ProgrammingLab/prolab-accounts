package profilestore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

type profileStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewProfileStore returns new profile store
func NewProfileStore(ctx context.Context, db *sql.DB) store.ProfileStore {
	return &profileStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *profileStoreImpl) CreateOrUpdateProfile(userID model.UserID, profile *record.Profile) (err error) {
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

	if profile.ID == 0 {
		err = profile.Insert(s.ctx, tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return errors.WithStack(err)
		}
		u := record.User{
			ID:        int64(userID),
			ProfileID: null.Int64From(profile.ID),
		}
		_, err = u.Update(s.ctx, tx, boil.Whitelist("profile_id", "updated_at"))
		if err != nil {
			_ = tx.Rollback()
			return errors.WithStack(err)
		}
	} else {
		_, err = profile.Update(s.ctx, tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return errors.WithStack(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}
	return nil
}
