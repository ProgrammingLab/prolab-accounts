package githubstore

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type githubStoreImpl struct {
	ctx context.Context
	db  *sqlutil.DB
	cli *redis.Client
}

// NewGitHubStore returns new github store
func NewGitHubStore(ctx context.Context, db *sqlutil.DB, cli *redis.Client) store.GitHubStore {
	return &githubStoreImpl{
		ctx: ctx,
		db:  db,
		cli: cli,
	}
}

const (
	contributionTotalCount = "contributions-total-count"
)

func (s *githubStoreImpl) UpdateContributionDays(c *model.GitHubContributionCollection) ([]*record.GithubContributionDay, error) {
	z := redis.Z{
		Score:  float64(c.TotalCount),
		Member: int64(c.UserID),
	}
	err := s.cli.ZAdd(contributionTotalCount, z).Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	days := c.Days
	err = s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		q := record.GithubContributionDayWhere.UserID.EQ(int64(c.UserID))
		_, err := record.GithubContributionDays(q).DeleteAll(ctx, tx)
		if err != nil {
			return errors.WithStack(err)
		}

		return bulkInsert(ctx, tx, days)
	})

	if err != nil {
		return nil, err
	}

	return days, nil
}

func bulkInsert(ctx context.Context, tx *sql.Tx, days []*record.GithubContributionDay) error {
	if len(days) == 0 {
		return nil
	}

	q := &strings.Builder{}
	_, err := q.WriteString("INSERT INTO " + record.TableNames.GithubContributionDays +
		" (count, date, user_id, created_at, updated_at) VALUES ")
	if err != nil {
		return errors.WithStack(err)
	}

	verbs, v := insertValueQuery(1)
	_, err = q.WriteString(v)
	if err != nil {
		return errors.WithStack(err)
	}
	for i := 0; i < len(days)-1; i++ {
		nxt, v := insertValueQuery(verbs)
		verbs = nxt
		_, err := q.WriteString(", " + v)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	_, err = q.WriteString(";")
	if err != nil {
		return errors.WithStack(err)
	}

	stmt, err := tx.Prepare(q.String())
	if err != nil {
		return errors.WithStack(err)
	}

	values := make([]interface{}, 0, len(days)*5)
	now := time.Now().In(boil.GetLocation())
	for _, d := range days {
		values = append(values, insertValues(d, now, now)...)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, q.String())
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = stmt.ExecContext(ctx, values...)
	return errors.WithStack(err)
}

func insertValueQuery(verbs int) (int, string) {
	return verbs + 5, fmt.Sprintf("($%v, $%v, $%v, $%v, $%v)", verbs, verbs+1, verbs+2, verbs+3, verbs+4)
}

func insertValues(d *record.GithubContributionDay, createdAt, updatedAt time.Time) []interface{} {
	return []interface{}{
		d.Count,
		d.Date,
		d.UserID,
		createdAt,
		updatedAt,
	}
}
