package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// UserStore accesses users data
type UserStore interface {
	CreateUser(user *record.User) error
	GetPublicUserByName(name string) (*record.User, error)
	GetUserByName(name string) (*record.User, error)
	GetUserByEmail(email string) (*record.User, error)
	GetUserWithPrivate(userID model.UserID) (*record.User, error)
	ListPublicUsers(subName string, minUserID model.UserID, limit int) ([]*record.User, model.UserID, error)
	ListPrivateUsers(subName string, minUserID model.UserID, limit int) ([]*record.User, model.UserID, error)
	UpdateFullName(userID model.UserID, fullName string) (*record.User, error)
	UpdateIcon(userID model.UserID, filename string) (u *record.User, old string, err error)
}
