package invitationstore

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base32"
	"strings"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
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

const (
	invitationCodeBytes = 32
)

func (s *invitationStoreImpl) ListInvitations() ([]*record.Invitation, error) {
	mods := []qm.QueryMod{
		qm.OrderBy(record.InvitationColumns.CreatedAt),
	}

	invs, err := record.Invitations(mods...).All(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return invs, nil
}

func (s *invitationStoreImpl) GetInvitation(id int64) (*record.Invitation, error) {
	inv, err := record.FindInvitation(s.ctx, s.db, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return inv, nil
}

func (s *invitationStoreImpl) CreateInvitation(inviter model.UserID, email string) (*record.Invitation, error) {
	code, err := generateInvitationCode()
	if err != nil {
		return nil, err
	}

	_, err = record.Invitations(record.InvitationWhere.Email.EQ(email)).DeleteAll(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	inv := &record.Invitation{
		Code:      code,
		Email:     email,
		InviterID: int64(inviter),
	}
	err = inv.Insert(s.ctx, s.db, boil.Infer())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return inv, nil
}

func (s *invitationStoreImpl) DeleteInvitation(id int64) error {
	_, err := record.Invitations(record.InvitationWhere.ID.EQ(id)).DeleteAll(s.ctx, s.db)
	return errors.WithStack(err)
}

func generateInvitationCode() (string, error) {
	b := make([]byte, invitationCodeBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.WithStack(err)
	}

	enc := base32.StdEncoding.WithPadding(base32.NoPadding)
	return strings.ToLower(enc.EncodeToString(b)), nil
}
