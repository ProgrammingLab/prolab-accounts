package rolestore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
)

type roleStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewRoleStore returns new role store
func NewRoleStore(ctx context.Context, db *sql.DB) store.RoleStore {
	return &roleStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *roleStoreImpl) ListRoles() ([]*record.Role, error) {
	roles, err := record.Roles().All(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return roles, nil
}

func (s *roleStoreImpl) GetRole(roleID int64) (*record.Role, error) {
	r, err := record.FindRole(s.ctx, s.db, roleID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return r, nil
}
