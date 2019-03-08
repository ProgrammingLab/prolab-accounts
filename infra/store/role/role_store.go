package rolestore

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type roleStoreImpl struct {
	ctx context.Context
	db  *sqlutil.DB
}

// NewRoleStore returns new role store
func NewRoleStore(ctx context.Context, db *sqlutil.DB) store.RoleStore {
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
