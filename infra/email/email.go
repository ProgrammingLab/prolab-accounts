package email

import (
	"context"

	"github.com/ProgrammingLab/prolab-accounts/static"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// Sender represents interface of email sender
type Sender interface {
	SendInvitationEmail(email string, req *record.Invitation) error
}

// NewSender creates new sender
func NewSender(ctx context.Context, cfg *config.Config, asset *static.EmailAsset) Sender {
	return &senderImpl{
		ctx:   ctx,
		cfg:   cfg,
		asset: asset,
	}
}

type senderImpl struct {
	ctx   context.Context
	cfg   *config.Config
	asset *static.EmailAsset
}

// SendInvitationEmail sends invitation email
func (s *senderImpl) SendInvitationEmail(email string, inv *record.Invitation) error {
	return nil
}
