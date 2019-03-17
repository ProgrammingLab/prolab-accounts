package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// PasswordResetStore provides password resets
type PasswordResetStore interface {
	CreateConfirmation(userID model.UserID, email string) (r *record.PasswordReset, token string, err error)
	GetConfirmation(email, token string) (*record.PasswordReset, error)
	UpdatePassword(reset *record.PasswordReset, password string) error
}
