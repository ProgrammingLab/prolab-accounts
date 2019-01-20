package server

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/labstack/gommon/log"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/grpclog"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
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

func (s *oAuthServiceServerImpl) StartOauthLogin(ctx context.Context, req *api_pb.StartOauthLoginRequest) (*api_pb.StartOAuthLoginResponse, error) {
	cli := s.HydraClient(ctx)
	challenge := req.GetLoginChallenge()
	res, _, err := cli.GetLoginRequest(challenge)
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}
	if res.Skip {
		res, _, err := cli.AcceptLoginRequest(challenge, swagger.AcceptLoginRequest{})
		if err != nil {
			grpclog.Error(err)
			return nil, err
		}

		resp := &api_pb.StartOAuthLoginResponse{
			Skip:        true,
			RedirectUrl: res.RedirectTo,
		}
		return resp, nil
	}

	resp := &api_pb.StartOAuthLoginResponse{
		Skip: false,
	}
	return resp, nil
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
	res, _, err := cli.AcceptLoginRequest(req.GetChallenge(), acReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp := &api_pb.OAuthLoginResponse{
		RedirectUrl: res.RedirectTo,
	}

	return resp, nil
}
