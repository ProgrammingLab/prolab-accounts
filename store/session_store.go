package store

import "github.com/ProgrammingLab/prolab-accounts/model"

// SessionStore accesses sessions
type SessionStore interface {
	CreateSession(userID model.UserID) (*model.Session, error)
}
