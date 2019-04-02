package store

import "github.com/ProgrammingLab/prolab-accounts/model"

// SessionStore accesses sessions
type SessionStore interface {
	CreateSession(nameOrEmail, password string) (*model.Session, error)
	GetSession(sessionID string) (*model.Session, error)
	ResetSession(userID model.UserID) (*model.Session, error)
	DeleteSession(sessionID string) error
}
