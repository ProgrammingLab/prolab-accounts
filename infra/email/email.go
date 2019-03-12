package email

import (
	"bytes"
	"context"
	"text/template"

	"github.com/jordan-wright/email"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/static"
)

// Sender represents interface of email sender
type Sender interface {
	SendInvitationEmail(req *record.Invitation) error
	SendEmailConfirmation(conf *record.EmailConfirmation) error
	SendEmailChanged(user *record.User, oldEmail string) error
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

const (
	subjectPrefix = "[プロラボアカウント]"
)

type invitationEmailData struct {
	RegistrationURL string
}

// SendInvitationEmail sends invitation email
func (s *senderImpl) SendInvitationEmail(inv *record.Invitation) error {
	tmpl, err := s.asset.GetTemplate("invitation.tmpl")
	if err != nil {
		return err
	}

	d := invitationEmailData{
		RegistrationURL: s.cfg.ClientRegistrationURL + "/" + inv.Code,
	}
	return s.send(inv.Email, "ユーザー登録", tmpl, d)
}

type emailConfirmationData struct {
	Name            string
	Email           string
	ConfirmationURL string
}

// SendEmailConfirmation sends email confirmation
func (s *senderImpl) SendEmailConfirmation(conf *record.EmailConfirmation) error {
	tmpl, err := s.asset.GetTemplate("email_confirmation.tmpl")
	if err != nil {
		return err
	}

	d := emailConfirmationData{
		Name:            conf.R.User.Name,
		Email:           conf.Email,
		ConfirmationURL: s.cfg.ClientConfirmationURL + "/" + conf.Token,
	}
	return s.send(conf.Email, "メールアドレスの確認", tmpl, d)
}

type emailChangedData struct {
	Name  string
	Email string
}

func (s *senderImpl) SendEmailChanged(user *record.User, oldEmail string) error {
	tmpl, err := s.asset.GetTemplate("email_changed.tmpl")
	if err != nil {
		return err
	}

	d := emailChangedData{
		Name:  user.Name,
		Email: user.Email,
	}
	return s.send(oldEmail, "メールアドレスが変更されました", tmpl, d)
}

func (s *senderImpl) send(to, subject string, tmpl *template.Template, d interface{}) error {
	buf := &bytes.Buffer{}
	err := tmpl.Execute(buf, d)
	if err != nil {
		return errors.WithStack(err)
	}

	e := email.NewEmail()
	e.From = s.cfg.EmailFrom
	e.To = []string{to}
	e.Subject = subjectPrefix + subject
	e.Text = buf.Bytes()
	err = e.Send(s.cfg.SMTPAddr, nil)
	return errors.WithStack(err)
}
