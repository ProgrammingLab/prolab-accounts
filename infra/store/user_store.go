package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// UserStore accesses users data
type UserStore interface {
	CreateUser(user *record.User) error
	GetUser(userID model.UserID) (*record.User, error)
}
