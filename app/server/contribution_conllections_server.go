package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

// ContributionConllectionServiceServer is a composite interface of api_pb.ContributionConllectionServiceServer and grapiserver.Server.
type ContributionConllectionServiceServer interface {
	api_pb.ContributionConllectionServiceServer
	grapiserver.Server
}

// NewContributionConllectionServiceServer creates a new ContributionConllectionServiceServer instance.
func NewContributionConllectionServiceServer(store di.StoreComponent) ContributionConllectionServiceServer {
	return &contributionConllectionServiceServerImpl{
		StoreComponent: store,
	}
}

type contributionConllectionServiceServerImpl struct {
	di.StoreComponent
}

func (s *contributionConllectionServiceServerImpl) ListContributionConllections(ctx context.Context, req *api_pb.ListContributionConllectionsRequest) (*api_pb.ListContributionConllectionsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
