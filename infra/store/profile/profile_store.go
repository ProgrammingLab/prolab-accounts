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

func (s *profileStoreImpl) CreateOrUpdateProfile(userID model.UserID, profile *record.Profile, updateRole bool) error {
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
			cols := []string{"updated_at", "description", "grade", "left", "department_id", "twitter_screen_name", "github_user_name", "profile_scope", "display_name", "atcoder_user_name"}
			if updateRole {
				cols = append(cols, "role_id")
			}
			li := boil.Whitelist(cols...)
			_, err := profile.Update(s.ctx, tx, li)
			if err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})

	return err
}
