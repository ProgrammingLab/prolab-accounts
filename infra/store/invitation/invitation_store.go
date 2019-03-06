package invitationstore

import (
	"context"
	"database/sql"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
)

type invitationStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewInvitationStore returns new invitation store
func NewInvitationStore(ctx context.Context, db *sql.DB) store.InvitationStore {
	return &invitationStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *invitationStoreImpl) ListInvitations() ([]*record.Invitation, error) {
	panic("not implemented")
}

func (s *invitationStoreImpl) GetInvitation(id int64) (*record.Invitation, error) {
	panic("not implemented")
}

func (s *invitationStoreImpl) CreateInvitation(email string) (*record.Invitation, error) {
	panic("not implemented")
}

func (s *invitationStoreImpl) DeleteInvitation(id int64) error {
	panic("not implemented")
}
