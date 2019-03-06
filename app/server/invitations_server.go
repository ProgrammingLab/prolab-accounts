package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
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
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *invitationServiceServerImpl) DeleteInvitation(ctx context.Context, req *api_pb.DeleteInvitationRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
