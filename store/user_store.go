package store

import "github.com/ProgrammingLab/prolab-accounts/dao"

// UserStore accesses users data
type UserStore interface {
	FindUserByEmailOrName(name string) (*dao.User, error)
}
