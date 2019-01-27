package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/labstack/gommon/log"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	type_pb "github.com/ProgrammingLab/prolab-accounts/api/type"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
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
	res, resp, err := cli.GetLoginRequest(challenge)
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}
	if err := hydraError(resp); err != nil {
		grpclog.Error(err)
		return nil, err
	}

	if res.Skip {
		res, resp, err := cli.AcceptLoginRequest(challenge, swagger.AcceptLoginRequest{})
		if err != nil {
			grpclog.Error(err)
			return nil, err
		}
		if err := hydraError(resp); err != nil {
			grpclog.Error(err)
			return nil, err
		}

		return &api_pb.StartOAuthLoginResponse{
			Skip:        true,
			RedirectUrl: res.RedirectTo,
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
		log.Error(err)
		return nil, util.ErrInternalServer
	}

	cli := s.HydraClient(ctx)
	acReq := swagger.AcceptLoginRequest{
		Subject:     strconv.FormatInt(int64(session.UserID), 10),
		Remember:    req.Remember,
		RememberFor: int64(time.Hour.Seconds()),
	}
	res, resp, err := cli.AcceptLoginRequest(req.GetLoginChallenge(), acReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err := hydraError(resp); err != nil {
		grpclog.Error(err)
		return nil, err
	}

	return &api_pb.OAuthLoginResponse{
		RedirectUrl: res.RedirectTo,
	}, nil
}

func (s *oAuthServiceServerImpl) StartOAuthConsent(ctx context.Context, req *api_pb.StartOAuthConsentRequest) (*api_pb.StartOAuthConsentResponse, error) {
	cli := s.HydraClient(ctx)
	challenge := req.GetConsentChallenge()
	res, resp, err := cli.GetConsentRequest(challenge)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err := hydraError(resp); err != nil {
		grpclog.Error(err)
		return nil, err
	}

	if res.Skip {
		req := swagger.AcceptConsentRequest{
			GrantScope:               res.RequestedScope,
			GrantAccessTokenAudience: res.RequestedAccessTokenAudience,
		}
		res, resp, err := cli.AcceptConsentRequest(challenge, req)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if err := hydraError(resp); err != nil {
			grpclog.Error(err)
			return nil, err
		}

		return &api_pb.StartOAuthConsentResponse{
			Skip:        true,
			RedirectUrl: res.RedirectTo,
		}, nil
	}

	return &api_pb.StartOAuthConsentResponse{
		Skip:            false,
		RequestedScopes: res.RequestedScope,
		Client:          clientToResponse(res.Client),
	}, nil
}

func (s *oAuthServiceServerImpl) OAuthConsent(ctx context.Context, req *api_pb.OAuthConsentRequest) (*api_pb.OAuthConsentResponse, error) {
	challenge := req.GetConsentChallenge()
	cli := s.HydraClient(ctx)
	if req.GetAccept() {
		cons, resp, err := cli.GetConsentRequest(challenge)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if err := hydraError(resp); err != nil {
			grpclog.Error(err)
			return nil, err
		}

		acReq := swagger.AcceptConsentRequest{
			GrantScope:               req.GetGrantScopes(),
			GrantAccessTokenAudience: cons.RequestedAccessTokenAudience,
			Remember:                 req.GetRemember(),
			RememberFor:              int64(time.Hour.Seconds()),
		}
		res, resp, err := cli.AcceptConsentRequest(challenge, acReq)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if err := hydraError(resp); err != nil {
			grpclog.Error(err)
			return nil, err
		}

		return &api_pb.OAuthConsentResponse{
			RedirectUrl: res.RedirectTo,
		}, nil
	}

	rej := swagger.RejectRequest{
		Error_:           "access_denied",
		ErrorDescription: "The resource owner denied the request",
	}
	res, resp, err := cli.RejectConsentRequest(challenge, rej)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err := hydraError(resp); err != nil {
		grpclog.Error(err)
		return nil, err
	}

	return &api_pb.OAuthConsentResponse{
		RedirectUrl: res.RedirectTo,
	}, nil
}

func clientToResponse(cli swagger.OAuth2Client) *type_pb.Client {
	return &type_pb.Client{
		Id:       cli.ClientId,
		Name:     cli.ClientName,
		Uri:      cli.ClientUri,
		Contacts: cli.Contacts,
		LogoUri:  cli.LogoUri,
		Owner:    cli.Owner,
	}
}

func hydraError(resp *swagger.APIResponse) error {
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	generic := &swagger.GenericError{}
	err := json.Unmarshal(resp.Payload, generic)
	if err != nil {
		return errors.WithStack(err)
	}
	return status.Error(util.CodeFromHTTPStatus(resp.StatusCode), generic.Error_)
}
