package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

// OAuthServiceServer is a composite interface of api_pb.OAuthServiceServer and grapiserver.Server.
type OAuthServiceServer interface {
	api_pb.OAuthServiceServer
	grapiserver.Server
}

// NewOAuthServiceServer creates a new OAuthServiceServer instance.
func NewOAuthServiceServer(cli di.ClientComponent) OAuthServiceServer {
	return &oAuthServiceServerImpl{
		ClientComponent: cli,
	}
}

type oAuthServiceServerImpl struct {
	di.ClientComponent
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
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
