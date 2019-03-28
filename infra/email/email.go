package email

import (
	"bytes"
	"context"
	"crypto/tls"
	"net/url"

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
	SendPasswordReset(reset *record.PasswordReset, token string) error
	SendPasswordChanged(user *record.User) error
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
	d := invitationEmailData{
		RegistrationURL: s.cfg.ClientRegistrationURL + "/" + inv.Code,
	}
	return s.send(inv.Email, "ユーザー登録", "invitation.tmpl", d)
}

type emailConfirmationData struct {
	Name            string
	Email           string
	ConfirmationURL string
}

// SendEmailConfirmation sends email confirmation
func (s *senderImpl) SendEmailConfirmation(conf *record.EmailConfirmation) error {
	d := emailConfirmationData{
		Name:            conf.R.User.Name,
		Email:           conf.Email,
		ConfirmationURL: s.cfg.ClientConfirmationURL + "/" + conf.Token,
	}
	return s.send(conf.Email, "メールアドレスの確認", "email_confirmation.tmpl", d)
}

type emailChangedData struct {
	Name  string
	Email string
}

func (s *senderImpl) SendEmailChanged(user *record.User, oldEmail string) error {
	d := emailChangedData{
		Name:  user.Name,
		Email: user.Email,
	}
	return s.send(oldEmail, "メールアドレスが変更されました", "email_changed.tmpl", d)
}

type passwordResetData struct {
	Name            string
	ConfirmationURL string
}

func (s *senderImpl) SendPasswordReset(reset *record.PasswordReset, token string) error {
	u := reset.R.User
	v := url.Values{}
	v.Add("email", reset.Email)
	d := passwordResetData{
		Name:            u.Name,
		ConfirmationURL: s.cfg.ClientPasswordResetURL + "/" + token + "?" + v.Encode(),
	}
	return s.send(reset.Email, "パスワードのリセット", "password_reset.tmpl", d)
}

type passwordChangedData struct {
	Name string
}

func (s *senderImpl) SendPasswordChanged(user *record.User) error {
	d := passwordChangedData{
		Name: user.Name,
	}
	return s.send(user.Email, "パスワードが変更されました", "password_changed.tmpl", d)
}

func (s *senderImpl) send(to, subject, tmplName string, d interface{}) error {
	tmpl, err := s.asset.GetTemplate(tmplName)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, d)
	if err != nil {
		return errors.WithStack(err)
	}

	e := email.NewEmail()
	e.From = s.cfg.EmailFrom
	e.To = []string{to}
	e.Subject = subjectPrefix + subject
	e.Text = buf.Bytes()
	if s.cfg.SMTPInsecureSkipVerify {
		t := &tls.Config{
			InsecureSkipVerify: true,
		}
		err = e.SendWithTLS(s.cfg.SMTPAddr, nil, t)
		return errors.WithStack(err)
	}
	err = e.Send(s.cfg.SMTPAddr, nil)
	return errors.WithStack(err)
}
