package di

import (
	"database/sql"

	// for database/sql
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/store"
	userstore "github.com/ProgrammingLab/prolab-accounts/store/user"
)

// StoreComponent is an interface of stores
type StoreComponent interface {
	UserStore() store.UserStore
}

// NewStoreComponent returns new store component
func NewStoreComponent(cfg *config.Config) (StoreComponent, error) {
	db, err := sql.Open("mysql", cfg.DataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "faild to connect db")
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, errors.Wrap(err, "faild to ping db")
	}

	return &storeComponentImpl{db: db}, nil
}

type storeComponentImpl struct {
	db *sql.DB
}

func (s *storeComponentImpl) UserStore() store.UserStore {
	return userstore.NewUserStore(s.db)
}
