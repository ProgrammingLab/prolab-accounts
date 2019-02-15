package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/model"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// UserStore accesses users data
type UserStore interface {
	CreateUser(user *record.User) error
	GetUser(userID model.UserID) (*record.User, error)
}
