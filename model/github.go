package model

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// GitHubContributionCollection represents collection of github contribution
type GitHubContributionCollection struct {
	UserID     UserID
	TotalCount int
	Days       []*record.GithubContributionDay
}
