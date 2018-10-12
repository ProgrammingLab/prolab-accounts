package model

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/pkg/errors"
)

// Session represents a user session
type Session struct {
	ID     string
	UserID UserID
}

const (
	idLength = 32
)

// NewSession returns new session
func NewSession(userID UserID) (*Session, error) {
	id := make([]byte, idLength)
	_, err := rand.Read(id)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return &Session{
		ID:     hex.EncodeToString(id),
		UserID: userID,
	}, nil
}
