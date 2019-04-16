package achievementstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type achievementStoreImpl struct {
	ctx context.Context
	db  *sqlutil.DB
}

// NewAchievementStore creates new achievement store
func NewAchievementStore(ctx context.Context, db *sqlutil.DB) store.AchievementStore {
	return &achievementStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *achievementStoreImpl) CreateAchievement(ach *record.Achievement, memberIDs []model.UserID) (*record.Achievement, error) {
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		ach.ID = 0
		err := ach.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return errors.WithStack(err)
		}

		members := make([]*record.AchievementUser, 0, len(memberIDs))
		for i, id := range memberIDs {
			members = append(members, &record.AchievementUser{
				UserID:   int64(id),
				Priority: i,
			})
		}

		err = ach.AddAchievementUsers(ctx, tx, true, members...)
		if err != nil {
			return errors.WithStack(err)
		}

		mods := []qm.QueryMod{
			s.membersOrder(),
			qm.Load("User"),
		}
		err = ach.L.LoadAchievementUsers(ctx, tx, true, ach, sqlutil.QueryMods(mods))
		return errors.WithStack(err)
	})
	if err != nil {
		return nil, err
	}

	ach, err = s.GetAchievement(ach.ID)
	return ach, err
}

func (s *achievementStoreImpl) GetAchievement(id int64) (*record.Achievement, error) {
	mods := []qm.QueryMod{
		record.AchievementWhere.ID.EQ(id),
		record.AchievementWhere.DeletedAt.IsNull(),
	}
	mods = append(mods, s.load()...)
	ach, err := record.Achievements(mods...).One(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ach, nil
}

func (s *achievementStoreImpl) ListAchievements(before time.Time, limit int) (aches []*record.Achievement, next time.Time, err error) {
	mods := []qm.QueryMod{
		record.AchievementWhere.HappenedAt.LT(before),
		record.AchievementWhere.DeletedAt.IsNull(),
		qm.OrderBy(record.AchievementColumns.HappenedAt + " desc"),
	}
	mods = append(mods, s.load()...)
	aches, err = record.Achievements(mods...).All(s.ctx, s.db)
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	if len(aches) == 0 {
		return aches, time.Time{}, nil
	}
	return aches, aches[len(aches)-1].HappenedAt, nil
}

func (s *achievementStoreImpl) UpdateAchievement(ach *record.Achievement, memberIDs []model.UserID) (*record.Achievement, error) {
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, err := ach.Update(ctx, tx, boil.Whitelist("title", "award", "url", "description", "happened_at", "updated_at"))
		if err != nil {
			return errors.WithStack(err)
		}

		members := make([]*record.AchievementUser, 0, len(memberIDs))
		for i, id := range memberIDs {
			members = append(members, &record.AchievementUser{
				UserID:   int64(id),
				Priority: i,
			})
		}

		_, err = ach.AchievementUsers().DeleteAll(ctx, tx)
		if err != nil {
			return errors.WithStack(err)
		}

		err = ach.AddAchievementUsers(ctx, tx, true, members...)
		return errors.WithStack(err)
	})
	if err != nil {
		return nil, err
	}

	ach, err = s.GetAchievement(ach.ID)
	return ach, err
}

func (s *achievementStoreImpl) UpdateAchievementImage(id int64, filename string) (ach *record.Achievement, old string, err error) {
	err = s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		ach, err = record.Achievements(qm.Where("id = ?", id), qm.Load("AchievementUsers.User")).One(ctx, tx)
		if err != nil {
			return errors.WithStack(err)
		}

		old = ach.ImageFilename.String

		ach.ImageFilename = null.StringFrom(filename)
		_, err = ach.Update(ctx, tx, boil.Whitelist("image_filename", "updated_at"))
		return errors.WithStack(err)
	})

	return ach, old, err
}

func (s *achievementStoreImpl) DeleteAchievement(id int64) error {
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		ach, err := record.FindAchievement(ctx, tx, id)
		if err != nil {
			return errors.WithStack(err)
		}

		ach.DeletedAt = null.TimeFrom(time.Now().In(boil.GetLocation()))
		_, err = ach.Update(ctx, tx, boil.Whitelist("deleted_at", "updated_at"))
		return errors.WithStack(err)
	})
	return err
}

func (s *achievementStoreImpl) load() []qm.QueryMod {
	return []qm.QueryMod{
		qm.Load("AchievementUsers", s.membersOrder()),
		qm.Load("AchievementUsers.User.Profile"),
	}
}

func (s *achievementStoreImpl) membersOrder() qm.QueryMod {
	return qm.OrderBy(record.AchievementUserColumns.Priority)
}
