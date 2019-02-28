package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// DepartmentStore accesses department data
type DepartmentStore interface {
	ListDepartments() ([]*record.Department, error)
	GetDepartment(roleID int64) (*record.Department, error)
}
