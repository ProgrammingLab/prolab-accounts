package store

import (
	"github.com/ProgrammingLab/prolab-accounts/dao"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// UserStore accesses users data
type UserStore interface {
	CreateUser(user *dao.User) error
	GetUser(userID model.UserID) (*dao.User, error)
	FindUserByEmailOrName(name string) (*dao.User, error)
}
