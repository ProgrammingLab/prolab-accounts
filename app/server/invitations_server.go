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
	_, err := s.getAdmin(ctx)
	if err != nil {
		return nil, err
	}

	is := s.InvitationStore(ctx)
	invs, err := is.ListInvitations()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp := &api_pb.ListInvitationsResponse{
		Invitations: invitationsToResponse(invs),
	}
	return resp, nil
}

func (s *invitationServiceServerImpl) GetInvitation(ctx context.Context, req *api_pb.GetInvitationRequest) (*api_pb.Invitation, error) {
	is := s.InvitationStore(ctx)
	inv, err := is.GetInvitation(req.GetToken())
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return invitationToResponse(inv), nil
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
	_, err := s.getAdmin(ctx)
	if err != nil {
		return nil, err
	}

	is := s.InvitationStore(ctx)
	err = is.DeleteInvitation(int64(req.GetInvitationId()))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return &empty.Empty{}, nil
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

func invitationsToResponse(invs []*record.Invitation) []*api_pb.Invitation {
	resp := make([]*api_pb.Invitation, 0, len(invs))
	for _, i := range invs {
		resp = append(resp, invitationToResponse(i))
	}

	return resp
}

func invitationToResponse(inv *record.Invitation) *api_pb.Invitation {
	return &api_pb.Invitation{
		InvitationId: uint32(inv.ID),
		Email:        inv.Email,
	}
}
