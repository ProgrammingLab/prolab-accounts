package static

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/errors"
)

// GetEmailsBox returns email template rice box
func GetEmailsBox() (*rice.Box, error) {
	b, err := rice.FindBox("emails")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return b, nil
}
