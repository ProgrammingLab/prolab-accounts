package profilestore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type profileStoreImpl struct {
	ctx context.Context
	db  *sqlutil.DB
}

// NewProfileStore returns new profile store
func NewProfileStore(ctx context.Context, db *sqlutil.DB) store.ProfileStore {
	return &profileStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *profileStoreImpl) CreateOrUpdateProfile(userID model.UserID, profile *record.Profile) error {
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		if profile.ID == 0 {
			err := profile.Insert(s.ctx, tx, boil.Infer())
			if err != nil {
				return errors.WithStack(err)
			}
			u := record.User{
				ID:        int64(userID),
				ProfileID: null.Int64From(profile.ID),
			}
			_, err = u.Update(s.ctx, tx, boil.Whitelist("profile_id", "updated_at"))
			if err != nil {
				return errors.WithStack(err)
			}
		} else {
			_, err := profile.Update(s.ctx, tx, boil.Infer())
			if err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})

	return err
}
