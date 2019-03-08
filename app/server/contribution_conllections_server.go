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

// ContributionConllectionServiceServer is a composite interface of api_pb.ContributionConllectionServiceServer and grapiserver.Server.
type ContributionConllectionServiceServer interface {
	api_pb.ContributionConllectionServiceServer
	grapiserver.Server
}

// NewContributionConllectionServiceServer creates a new ContributionConllectionServiceServer instance.
func NewContributionConllectionServiceServer(store di.StoreComponent, cfg *config.Config) ContributionConllectionServiceServer {
	return &contributionConllectionServiceServerImpl{
		StoreComponent: store,
		cfg:            cfg,
	}
}

type contributionConllectionServiceServerImpl struct {
	di.StoreComponent
	cfg *config.Config
}

const (
	defaultContributionConllectionsLimit = 10
)

func (s *contributionConllectionServiceServerImpl) ListContributionConllections(ctx context.Context, req *api_pb.ListContributionConllectionsRequest) (*api_pb.ListContributionConllectionsResponse, error) {
	gs := s.GitHubStore(ctx)
	limit := req.GetUsersCount()
	if limit == 0 {
		limit = defaultContributionConllectionsLimit
	}
	cols, err := gs.ListContributionCollections(int(limit))
	if err != nil {
		return nil, err
	}

	resp := contributionConllectionsToResponse(cols, s.cfg)
	return &api_pb.ListContributionConllectionsResponse{
		ContributionConllections: resp,
	}, nil
}

func contributionConllectionsToResponse(cols []*model.GitHubContributionCollection, cfg *config.Config) []*api_pb.ContributionConllection {
	resp := make([]*api_pb.ContributionConllection, 0, len(cols))
	for _, c := range cols {
		rc := contributionConllectionToResponse(c, cfg)
		if rc == nil {
			continue
		}
		resp = append(resp, rc)
	}

	return resp
}

func contributionConllectionToResponse(col *model.GitHubContributionCollection, cfg *config.Config) *api_pb.ContributionConllection {
	if len(col.Days) == 0 {
		return nil
	}

	u := col.Days[0].R.User
	return &api_pb.ContributionConllection{
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
