package store

import (
	"errors"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

var (
	// ErrEmailAlreadyInUse is returned when email is already in use
	ErrEmailAlreadyInUse = errors.New("email is already in use")
)

// EmailConfirmationStore provides email confirmations
type EmailConfirmationStore interface {
	CreateConfirmation(userID model.UserID, email string) (*record.EmailConfirmation, error)
	GetConfirmation(token string) (*record.EmailConfirmation, error)
	ConfirmEmail(c *record.EmailConfirmation) (*record.User, error)
}
