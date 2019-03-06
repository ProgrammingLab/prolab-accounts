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
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// InvitationServiceServer is a composite interface of api_pb.InvitationServiceServer and grapiserver.Server.
type InvitationServiceServer interface {
	api_pb.InvitationServiceServer
	grapiserver.Server
}

// NewInvitationServiceServer creates a new InvitationServiceServer instance.
func NewInvitationServiceServer(store di.StoreComponent, cli di.ClientComponent) InvitationServiceServer {
	return &invitationServiceServerImpl{
		StoreComponent:  store,
		ClientComponent: cli,
	}
}

var (
	// ErrEmailAlreadyInUse is returned when email is already in use
	ErrEmailAlreadyInUse = status.Error(codes.AlreadyExists, "email is already in use")
)

type invitationServiceServerImpl struct {
	di.StoreComponent
	di.ClientComponent
}

func (s *invitationServiceServerImpl) ListInvitations(ctx context.Context, req *api_pb.ListInvitationsRequest) (*api_pb.ListInvitationsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *invitationServiceServerImpl) GetInvitation(ctx context.Context, req *api_pb.GetInvitationRequest) (*api_pb.Invitation, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *invitationServiceServerImpl) CreateInvitation(ctx context.Context, req *api_pb.CreateInvitationRequest) (*api_pb.Invitation, error) {
	admin, err := s.getAdmin(ctx)
	if err != nil {
		return nil, err
	}

	us := s.UserStore(ctx)
	_, err = us.GetUserByEmail(req.GetEmail())
	if err == nil {
		return nil, ErrEmailAlreadyInUse
	}
	if errors.Cause(err) != sql.ErrNoRows {
		return nil, err
	}

	is := s.InvitationStore(ctx)
	inv, err := is.CreateInvitation(model.UserID(admin.ID), req.GetEmail())
	if err != nil {
		return nil, err
	}

	sender := s.EmailSender(ctx)
	err = sender.SendInvitationEmail(inv)
	if err != nil {
		return nil, err
	}

	return invitationToResponse(inv), nil
}

func (s *invitationServiceServerImpl) DeleteInvitation(ctx context.Context, req *api_pb.DeleteInvitationRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *invitationServiceServerImpl) getAdmin(ctx context.Context) (*record.User, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	us := s.UserStore(ctx)
	u, err := us.GetUserWithPrivate(userID)
	if err != nil {
		return nil, err
	}

	if u.Authority != int(model.Admin) {
		return nil, util.ErrUnauthenticated
	}
	return u, nil
}

func invitationToResponse(inv *record.Invitation) *api_pb.Invitation {
	return &api_pb.Invitation{
		InvitationId: uint32(inv.ID),
		Email:        inv.Email,
	}
}
