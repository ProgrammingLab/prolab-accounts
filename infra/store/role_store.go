package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// RoleStore accesses roles
type RoleStore interface {
	ListRoles() ([]*record.Role, error)
	GetRole(roleID int64) (*record.Role, error)
}
