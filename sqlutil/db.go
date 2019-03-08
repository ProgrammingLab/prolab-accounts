package sqlutil

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
)

// DB is database with util
type DB struct {
	*sql.DB
}

// New creates new db
func New(db *sql.DB) *DB {
	return &DB{
		DB: db,
	}
}

// Watch calls f with transaction
func (d *DB) Watch(ctx context.Context, f func(ctx context.Context, tx *sql.Tx) error) (err error) {
	tx, err := d.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		if e := util.ErrorFromRecover(recover()); e != nil {
			err = e
			_ = tx.Rollback()
		}
	}()

	err = f(ctx, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return errors.WithStack(err)
	}

	return nil
}
