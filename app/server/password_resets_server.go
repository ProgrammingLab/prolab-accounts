package server

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// PasswordResetServiceServer is a composite interface of api_pb.PasswordResetServiceServer and grapiserver.Server.
type PasswordResetServiceServer interface {
	api_pb.PasswordResetServiceServer
	grapiserver.Server
}

// NewPasswordResetServiceServer creates a new PasswordResetServiceServer instance.
func NewPasswordResetServiceServer(store di.StoreComponent, cli di.ClientComponent, cfg *config.Config) PasswordResetServiceServer {
	return &passwordResetServiceServerImpl{
		StoreComponent:  store,
		ClientComponent: cli,
		cfg:             cfg,
	}
}

type passwordResetServiceServerImpl struct {
	di.StoreComponent
	di.ClientComponent
	cfg *config.Config
}

func (s *passwordResetServiceServerImpl) GetPasswordReset(ctx context.Context, req *api_pb.GetPasswordResetRequest) (*empty.Empty, error) {
	ps := s.PasswordResetStore(ctx)
	_, err := ps.GetConfirmation(req.GetEmail(), req.GetToken())
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *passwordResetServiceServerImpl) CreatePasswordReset(ctx context.Context, req *api_pb.CreatePasswordResetRequest) (*empty.Empty, error) {
	email := req.GetEmail()

	us := s.UserStore(ctx)
	u, err := us.GetUserByEmail(email)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return &empty.Empty{}, nil
		}
		return nil, err
	}

	ps := s.PasswordResetStore(ctx)
	p, token, err := ps.CreateConfirmation(model.UserID(u.ID), email)
	if err != nil {
		return nil, err
	}

	sender := s.EmailSender(ctx)
	err = sender.SendPasswordReset(p, token)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *passwordResetServiceServerImpl) UpdatePassword(ctx context.Context, req *api_pb.UpdatePasswordRequest) (*api_pb.Session, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
