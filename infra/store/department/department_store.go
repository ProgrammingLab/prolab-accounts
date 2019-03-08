package departmentstore

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type departmentStoreImpl struct {
	ctx context.Context
	db  *sqlutil.DB
}

// NewDepartmentStore returns new entry department store
func NewDepartmentStore(ctx context.Context, db *sqlutil.DB) store.DepartmentStore {
	return &departmentStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *departmentStoreImpl) ListDepartments() ([]*record.Department, error) {
	deps, err := record.Departments().All(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return deps, nil
}

func (s *departmentStoreImpl) GetDepartment(roleID int64) (*record.Department, error) {
	dep, err := record.FindDepartment(s.ctx, s.db, roleID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dep, nil
}
