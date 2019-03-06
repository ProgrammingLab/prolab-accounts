package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// InvitationStore provides invitations
type InvitationStore interface {
	ListInvitations() ([]*record.Invitation, error)
	GetInvitation(id int64) (*record.Invitation, error)
	CreateInvitation(email string) (*record.Invitation, error)
	DeleteInvitation(id int64) error
}
