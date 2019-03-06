package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// InvitationStore provides invitations
type InvitationStore interface {
	ListInvitations() ([]*record.Invitation, error)
	GetInvitation(id int64) (*record.Invitation, error)
	CreateInvitation(inviter model.UserID, email string) (*record.Invitation, error)
	DeleteInvitation(id int64) error
}
