package userstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
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

func (s *userStoreImpl) ListPublicUsers(minUserID model.UserID, limit int) ([]*record.User, model.UserID, error) {
	mods := []qm.QueryMod{
		qm.Load("Profile.Role"),
		qm.Load("Profile.Department"),
		qm.Select("profiles.*", s.selectQuery(model.Public)),
		qm.InnerJoin("profiles on profiles.id = users.profile_id"),
		qm.Where("? <= users.id", minUserID),
		qm.Where("profiles.profile_scope = ?", model.Public),
		qm.Limit(limit + 1),
		qm.OrderBy("users.id"),
	}

	u, err := record.Users(mods...).All(s.ctx, s.db)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	if len(u) <= limit {
		return u, 0, nil
	}
	return u[:limit], model.UserID(u[limit].ID), nil
}

func (s *userStoreImpl) UpdateFullName(userID model.UserID, fullName string) (u *record.User, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		if err = util.ErrorFromRecover(recover()); err != nil {
			_ = tx.Rollback()
		}
	}()
	u, err = record.FindUser(s.ctx, tx, int64(userID))
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	u.FullName = fullName
	_, err = u.Update(s.ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	return u, nil
}

var selectQuery = map[model.ProfileScope]string{
	model.MembersOnly: "users.id, users.name, users.full_name, users.avatar_filename, users.profile_id",
	model.Public:      "users.id, users.name, users.avatar_filename, users.profile_id",
	model.Private:     "users.*",
}

func (s *userStoreImpl) selectQuery(scope model.ProfileScope) string {
	q, ok := selectQuery[scope]
	if !ok {
		return selectQuery[model.Public]
	}
	return q
}
