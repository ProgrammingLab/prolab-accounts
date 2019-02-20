package profilestore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
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

func (s *profileStoreImpl) CreateOrUpdateProfile(profile *record.Profile) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if err = util.ErrorFromRecover(recover()); err != nil {
			_ = tx.Rollback()
		}
	}()

	err = profile.Upsert(s.ctx, tx, true, nil, boil.Infer(), boil.Infer())
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
