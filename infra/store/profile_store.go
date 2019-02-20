package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// ProfileStore accesses profiles data
type ProfileStore interface {
	CreateOrUpdateProfile(profile *record.Profile) error
}
