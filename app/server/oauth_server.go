package server

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/ory/hydra/sdk/go/hydra/client/admin"
	"github.com/ory/hydra/sdk/go/hydra/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	type_pb "github.com/ProgrammingLab/prolab-accounts/api/type"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

// OAuthServiceServer is a composite interface of api_pb.OAuthServiceServer and grapiserver.Server.
type OAuthServiceServer interface {
	api_pb.OAuthServiceServer
	grapiserver.Server
}

// NewOAuthServiceServer creates a new OAuthServiceServer instance.
func NewOAuthServiceServer(cli di.ClientComponent, store di.StoreComponent) OAuthServiceServer {
	return &oAuthServiceServerImpl{
		ClientComponent: cli,
		StoreComponent:  store,
	}
}

type oAuthServiceServerImpl struct {
	di.ClientComponent
	di.StoreComponent
}

func (s *oAuthServiceServerImpl) StartOAuthLogin(ctx context.Context, req *api_pb.StartOAuthLoginRequest) (*api_pb.StartOAuthLoginResponse, error) {
	cli := s.HydraClient(ctx)
	challenge := req.GetLoginChallenge()
	params := &admin.GetLoginRequestParams{
		Context:        ctx,
		LoginChallenge: challenge,
	}
	res, err := cli.Admin.GetLoginRequest(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if res.Payload.Skip {
		params := &admin.AcceptLoginRequestParams{
			Context:        ctx,
			LoginChallenge: challenge,
		}
		res, err := cli.Admin.AcceptLoginRequest(params)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return &api_pb.StartOAuthLoginResponse{
			Skip:        true,
			RedirectUrl: res.Payload.RedirectTo,
		}, nil
	}

	return &api_pb.StartOAuthLoginResponse{
		Skip: false,
	}, nil
}

func (s *oAuthServiceServerImpl) OAuthLogin(ctx context.Context, req *api_pb.OAuthLoginRequest) (*api_pb.OAuthLoginResponse, error) {
	ss := s.SessionStore(ctx)
	session, err := ss.CreateSession(req.GetName(), req.GetPassword())
	if err != nil {
		if c := errors.Cause(err); c == sql.ErrNoRows || c == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errLogin
		}
		return nil, err
	}

	cli := s.HydraClient(ctx)
	sub := strconv.FormatInt(int64(session.UserID), 10)
	params := &admin.AcceptLoginRequestParams{
		Context:        ctx,
		LoginChallenge: req.GetLoginChallenge(),
		Body: &models.HandledLoginRequest{
			Subject:     &sub,
			Remember:    req.Remember,
			RememberFor: int64(time.Hour.Seconds()),
		},
	}
	res, err := cli.Admin.AcceptLoginRequest(params)
	if err != nil {
		return nil, err
	}

	return &api_pb.OAuthLoginResponse{
		RedirectUrl: res.Payload.RedirectTo,
	}, nil
}

func (s *oAuthServiceServerImpl) StartOAuthConsent(ctx context.Context, req *api_pb.StartOAuthConsentRequest) (*api_pb.StartOAuthConsentResponse, error) {
	cli := s.HydraClient(ctx)
	challenge := req.GetConsentChallenge()
	params := &admin.GetConsentRequestParams{
		Context:          ctx,
		ConsentChallenge: challenge,
	}
	res, err := cli.Admin.GetConsentRequest(params)
	if err != nil {
		return nil, err
	}

	if res.Payload.Skip {
		body := &models.HandledConsentRequest{
			GrantedScope:    res.Payload.RequestedScope,
			GrantedAudience: res.Payload.RequestedAudience,
		}
		params := &admin.AcceptConsentRequestParams{
			Context:          ctx,
			ConsentChallenge: challenge,
			Body:             body,
		}
		res, err := cli.Admin.AcceptConsentRequest(params)
		if err != nil {
			return nil, err
		}

		return &api_pb.StartOAuthConsentResponse{
			Skip:        true,
			RedirectUrl: res.Payload.RedirectTo,
		}, nil
	}

	return &api_pb.StartOAuthConsentResponse{
		Skip:            false,
		RequestedScopes: res.Payload.RequestedScope,
		Client:          clientToResponse(res.Payload.Client),
	}, nil
}

func (s *oAuthServiceServerImpl) OAuthConsent(ctx context.Context, req *api_pb.OAuthConsentRequest) (*api_pb.OAuthConsentResponse, error) {
	challenge := req.GetConsentChallenge()
	cli := s.HydraClient(ctx)
	if req.GetAccept() {
		params := &admin.GetConsentRequestParams{
			Context:          ctx,
			ConsentChallenge: challenge,
		}
		cons, err := cli.Admin.GetConsentRequest(params)
		if err != nil {
			return nil, err
		}

		body := &models.HandledConsentRequest{
			GrantedScope:    req.GetGrantScopes(),
			GrantedAudience: cons.Payload.RequestedAudience,
			Remember:        req.GetRemember(),
			RememberFor:     int64(time.Hour.Seconds()),
		}
		acParams := &admin.AcceptConsentRequestParams{
			Context:          ctx,
			ConsentChallenge: challenge,
			Body:             body,
		}
		res, err := cli.Admin.AcceptConsentRequest(acParams)
		if err != nil {
			return nil, err
		}

		return &api_pb.OAuthConsentResponse{
			RedirectUrl: res.Payload.RedirectTo,
		}, nil
	}

	body := &models.RequestDeniedError{
		Name:        "access_denied",
		Description: "The resource owner denied the request",
	}
	params := &admin.RejectConsentRequestParams{
		Context:          ctx,
		Body:             body,
		ConsentChallenge: challenge,
	}
	res, err := cli.Admin.RejectConsentRequest(params)
	if err != nil {
		return nil, err
	}

	return &api_pb.OAuthConsentResponse{
		RedirectUrl: res.Payload.RedirectTo,
	}, nil
}

func clientToResponse(cli *models.Client) *type_pb.Client {
	return &type_pb.Client{
		Id:       cli.ClientID,
		Name:     cli.Name,
		Uri:      cli.ClientURI,
		Contacts: cli.Contacts,
		LogoUri:  cli.LogoURI,
		Owner:    cli.Owner,
	}
}
