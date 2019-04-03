package store

import (
	"time"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// AchievementStore provides achievements
type AchievementStore interface {
	CreateAchievement(ach *record.Achievement, memberIDs []model.UserID) (*record.Achievement, error)
	GetAchievement(id int64) (*record.Achievement, error)
	ListAchievements(before time.Time, limit int) (aches []*record.Achievement, next time.Time, err error)
	UpdateAchievement(ach *record.Achievement, memberIDs []model.UserID) (*record.Achievement, error)
	UpdateAchievementImage(id int64, filename string) (ach *record.Achievement, old string, err error)
	DeleteAchievement(id int64) error
}
