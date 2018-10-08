package userstore

import (
	"database/sql"

	"github.com/ProgrammingLab/prolab-accounts/store"
)

type userStoreImpl struct {
	db *sql.DB
}

// NewUserStore returns new user store
func NewUserStore(db *sql.DB) store.UserStore {
	return &userStoreImpl{db: db}
}
