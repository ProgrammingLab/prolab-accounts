package model

import (
	"crypto/rand"
	"encoding/base32"
	"strings"

	"github.com/pkg/errors"
)

// GenerateSecureToken generates random token
func GenerateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.WithStack(err)
	}

	enc := base32.StdEncoding.WithPadding(base32.NoPadding)
	t := enc.EncodeToString(b)
	return strings.ToLower(t), nil
}
