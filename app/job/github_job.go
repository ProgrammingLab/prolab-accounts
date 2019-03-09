package job

import (
	"context"
	"time"

	"google.golang.org/grpc/grpclog"

	"github.com/pkg/errors"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

type githubUser struct {
	User struct {
		ContributionsCollection struct {
			ContributionCalendar struct {
				TotalContributions githubv4.Int
				Weeks              []struct {
					ContributionDays []struct {
						ContributionCount githubv4.Int
						Date              githubv4.String
					}
				}
			}
		} `graphql:"contributionsCollection(from:$from, to:$to)"`
	} `graphql:"user(login:$login)"`
}

const (
	contributionsFromDay = 60
	githubDateFormat     = "2006-1-2"
)

func githubJob(ctx context.Context, store di.StoreComponent, cfg *config.Config) error {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHubAccessToken},
	)
	httpClient := oauth2.NewClient(ctx, src)
	cli := githubv4.NewClient(httpClient)

	us := store.UserStore(ctx)
	next := model.UserID(1)
	to := time.Now().UTC()
	from := to.AddDate(0, 0, -contributionsFromDay).Round(time.Hour * 24)
	for next != 0 {
		users, nxt, err := us.ListPublicUsers(next, 100)
		next = nxt
		if err != nil {
			return err
		}

		for _, u := range users {
			name := u.R.Profile.GithubUserName
			if name.IsZero() {
				continue
			}
			gu, err := getGitHubUser(ctx, cli, name.String, from, to)
			if err != nil {
				grpclog.Errorf("github job: %v", err)
				continue
			}

			err = storeGitHubContributions(ctx, store, u, gu)
			if err != nil {
				return err
			}

			time.Sleep(500 * time.Millisecond)
		}
	}

	return nil
}

func getGitHubUser(ctx context.Context, cli *githubv4.Client, name string, from time.Time, to time.Time) (*githubUser, error) {
	v := map[string]interface{}{
		"login": githubv4.String(name),
		"from":  githubv4.DateTime{Time: from},
		"to":    githubv4.DateTime{Time: to},
	}
	q := &githubUser{}
	err := cli.Query(ctx, q, v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return q, nil
}

func storeGitHubContributions(ctx context.Context, store di.StoreComponent, user *record.User, github *githubUser) error {
	days := make([]*record.GithubContributionDay, 0, contributionsFromDay)
	weeks := github.User.ContributionsCollection.ContributionCalendar.Weeks
	for _, w := range weeks {
		for _, d := range w.ContributionDays {
			date, err := time.Parse(githubDateFormat, string(d.Date))
			if err != nil {
				return errors.WithStack(err)
			}
			gd := &record.GithubContributionDay{
				Count:  int(d.ContributionCount),
				Date:   date,
				UserID: user.ID,
			}
			days = append(days, gd)
		}
	}

	c := &model.GitHubContributionCollection{
		UserID:     model.UserID(user.ID),
		TotalCount: int(github.User.ContributionsCollection.ContributionCalendar.TotalContributions),
		Days:       days,
	}
	gs := store.GitHubStore(ctx)
	_, err := gs.UpdateContributionDays(c)
	return err
}
