package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// ContributionCollectionServiceServer is a composite interface of api_pb.ContributionCollectionServiceServer and grapiserver.Server.
type ContributionCollectionServiceServer interface {
	api_pb.ContributionCollectionServiceServer
	grapiserver.Server
}

// NewContributionCollectionServiceServer creates a new ContributionCollectionServiceServer instance.
func NewContributionCollectionServiceServer(store di.StoreComponent, cfg *config.Config) ContributionCollectionServiceServer {
	return &contributionCollectionServiceServerImpl{
		StoreComponent: store,
		cfg:            cfg,
	}
}

type contributionCollectionServiceServerImpl struct {
	di.StoreComponent
	cfg *config.Config
}

const (
	defaultContributionCollectionsLimit = 10
)

func (s *contributionCollectionServiceServerImpl) ListContributionCollections(ctx context.Context, req *api_pb.ListContributionCollectionsRequest) (*api_pb.ListContributionCollectionsResponse, error) {
	gs := s.GitHubStore(ctx)
	limit := req.GetUsersCount()
	if limit == 0 {
		limit = defaultContributionCollectionsLimit
	}
	cols, err := gs.ListContributionCollections(int(limit))
	if err != nil {
		return nil, err
	}

	resp := contributionCollectionsToResponse(cols, s.cfg)
	return &api_pb.ListContributionCollectionsResponse{
		ContributionCollections: resp,
	}, nil
}

func contributionCollectionsToResponse(cols []*model.GitHubContributionCollection, cfg *config.Config) []*api_pb.ContributionCollection {
	resp := make([]*api_pb.ContributionCollection, 0, len(cols))
	for _, c := range cols {
		rc := contributionCollectionToResponse(c, cfg)
		if rc == nil {
			continue
		}
		resp = append(resp, rc)
	}

	return resp
}

func contributionCollectionToResponse(col *model.GitHubContributionCollection, cfg *config.Config) *api_pb.ContributionCollection {
	if len(col.Days) == 0 {
		return nil
	}

	u := col.Days[0].R.User
	return &api_pb.ContributionCollection{
		User:       userToResponse(u, false, cfg),
		TotalCount: int32(col.TotalCount),
		Days:       contributionDaysToResponse(col.Days),
	}
}

func contributionDaysToResponse(days []*record.GithubContributionDay) []*api_pb.ContributionDay {
	resp := make([]*api_pb.ContributionDay, 0, len(days))
	for _, d := range days {
		rd := &api_pb.ContributionDay{
			Date:  timeToResponse(d.Date),
			Count: int32(d.Count),
		}
		resp = append(resp, rd)
	}

	return resp
}
