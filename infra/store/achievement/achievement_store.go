package achievementstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
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
func (s *achievementStoreImpl) CreateAchievement(ach *record.Achievement, memberIDs []model.UserID) error {
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
		return errors.WithStack(err)
	})
	return err
}

func (s *achievementStoreImpl) ListAchievements(before time.Time, limit int) (aches []*record.Achievement, next time.Time, err error) {
	mods := []qm.QueryMod{
		qm.Load(record.AchievementRels.AchievementUsers, qm.OrderBy(record.AchievementUserColumns.Priority)),
		record.AchievementWhere.HappenedAt.LT(before),
		qm.OrderBy(record.AchievementColumns.HappenedAt),
	}
	aches, err = record.Achievements(mods...).All(s.ctx, s.db)
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	return aches, aches[len(aches)-1].HappenedAt, nil
}

func (s *achievementStoreImpl) UpdateAchievement(ach *record.Achievement, memberIDs []model.UserID) error {
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, err := ach.Update(ctx, tx, boil.Infer())
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
	return err
}

func (s *achievementStoreImpl) DeleteAchievement(id int64) error {
	_, err := record.Achievements(record.AchievementWhere.ID.EQ(id)).DeleteAll(s.ctx, s.db)
	return errors.WithStack(err)
}
