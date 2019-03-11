package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// EmailConfirmationServiceServer is a composite interface of api_pb.EmailConfirmationServiceServer and grapiserver.Server.
type EmailConfirmationServiceServer interface {
	api_pb.EmailConfirmationServiceServer
	grapiserver.Server
}

// NewEmailConfirmationServiceServer creates a new EmailConfirmationServiceServer instance.
func NewEmailConfirmationServiceServer() EmailConfirmationServiceServer {
	return &emailConfirmationServiceServerImpl{}
}

type emailConfirmationServiceServerImpl struct {
}

func (s *emailConfirmationServiceServerImpl) ListEmailConfirmations(ctx context.Context, req *api_pb.ListEmailConfirmationsRequest) (*api_pb.ListEmailConfirmationsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *emailConfirmationServiceServerImpl) GetEmailConfirmation(ctx context.Context, req *api_pb.GetEmailConfirmationRequest) (*api_pb.EmailConfirmation, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *emailConfirmationServiceServerImpl) CreateEmailConfirmation(ctx context.Context, req *api_pb.CreateEmailConfirmationRequest) (*api_pb.EmailConfirmation, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *emailConfirmationServiceServerImpl) UpdateEmailConfirmation(ctx context.Context, req *api_pb.UpdateEmailConfirmationRequest) (*api_pb.EmailConfirmation, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *emailConfirmationServiceServerImpl) DeleteEmailConfirmation(ctx context.Context, req *api_pb.DeleteEmailConfirmationRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
