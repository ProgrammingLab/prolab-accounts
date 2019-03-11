package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
)

// EmailConfirmationServiceServer is a composite interface of api_pb.EmailConfirmationServiceServer and grapiserver.Server.
type EmailConfirmationServiceServer interface {
	api_pb.EmailConfirmationServiceServer
	grapiserver.Server
}

// NewEmailConfirmationServiceServer creates a new EmailConfirmationServiceServer instance.
func NewEmailConfirmationServiceServer(store di.StoreComponent, cli di.ClientComponent, cfg *config.Config) EmailConfirmationServiceServer {
	return &emailConfirmationServiceServerImpl{
		StoreComponent:  store,
		ClientComponent: cli,
		cfg:             cfg,
	}
}

type emailConfirmationServiceServerImpl struct {
	di.StoreComponent
	di.ClientComponent
	cfg *config.Config
}

func (s *emailConfirmationServiceServerImpl) ConfirmEmail(ctx context.Context, req *api_pb.ConfirmEmailRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *emailConfirmationServiceServerImpl) CreateEmailConfirmation(ctx context.Context, req *api_pb.CreateEmailConfirmationRequest) (*empty.Empty, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	cs := s.EmailConfirmationStore(ctx)
	c, err := cs.CreateConfirmation(userID, req.GetNewEmail())
	if err != nil {
		if errors.Cause(err) == store.ErrEmailAlreadyInUse {
			return nil, ErrEmailAlreadyInUse
		}
		return nil, err
	}

	sender := s.EmailSender(ctx)
	err = sender.SendEmailConfirmation(c)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
