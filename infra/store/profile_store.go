package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// ProfileStore accesses profiles data
type ProfileStore interface {
	CreateOrUpdateProfile(userID model.UserID, profile *record.Profile, updateRole bool) error
}
