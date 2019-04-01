package store

import (
	"time"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// AchievementStore provides achievements
type AchievementStore interface {
	CreateAchievement(ach *record.Achievement, memberIDs []model.UserID) error
	GetAchievement(id int64) (*record.Achievement, error)
	ListAchievements(before time.Time, limit int) (aches []*record.Achievement, next time.Time, err error)
	UpdateAchievement(ach *record.Achievement, memberIDs []model.UserID) error
	DeleteAchievement(id int64) error
}
