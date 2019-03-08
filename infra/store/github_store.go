package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// GitHubStore provides github data
type GitHubStore interface {
	UpdateContributionDays(c *model.GitHubContributionCollection) ([]*record.GithubContributionDay, error)
	ListContributionCollections(usersLimit int) ([]*model.GitHubContributionCollection, error)
}
