package store

import (
	"time"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// AchievementStore provides achievements
type AchievementStore interface {
	CreateAchievement(ach *record.Achievement, memberIDs []int64) error
	ListAchievements(before time.Time, limit int) (aches []*record.Achievement, next time.Time, err error)
	UpdateAchievement(ach *record.Achievement, memberIDs []int64) error
	DeleteAchievement(id int64) error
}
