package entrystore

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

type entryStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewEntryStore returns new entry blog store
func NewEntryStore(ctx context.Context, db *sql.DB) store.EntryStore {
	return &entryStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *entryStoreImpl) ListPublicEntries(before time.Time, limit int) ([]*record.Entry, time.Time, error) {
	mods := []qm.QueryMod{
		qm.Load(record.EntryRels.Author),
		qm.Load(record.EntryRels.Blog),
		qm.InnerJoin("users on users.id = entries.author_id"),
		qm.InnerJoin("profiles on profiles.id = users.profile_id"),
		qm.Where("profiles.profile_scope = ?", model.Public),
		qm.Where("entries.published_at <= ?", before),
		qm.Limit(limit + 1),
		qm.OrderBy("entries.published_at desc"),
	}

	e, err := record.Entries(mods...).All(s.ctx, s.db)
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	if len(e) <= limit {
		return e, time.Unix(0, 0), nil
	}
	return e[:limit], e[limit].PublishedAt, nil
}

func (s *entryStoreImpl) CreateEntries(blog *record.Blog, feed *gofeed.Feed) (n int64, err error) {
	rev := make([]*gofeed.Item, len(feed.Items))
	for i, item := range feed.Items {
		rev[len(rev)-1-i] = item
	}
	feed.Items = rev

	tx, err := s.db.Begin()
	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer func() {
		if e := util.ErrorFromRecover(recover()); e != nil {
			_ = tx.Rollback()
			err = e
		}
	}()

	mods := []qm.QueryMod{
		qm.Select(record.EntryColumns.ID, record.EntryColumns.GUID),
		qm.Where("blog_id = ?", blog.ID),
	}
	entries, err := record.Entries(mods...).All(s.ctx, tx)
	if err != nil {
		_ = tx.Rollback()
		return 0, errors.WithStack(err)
	}

	exists := make(map[string]struct{})
	for _, e := range entries {
		exists[e.GUID] = struct{}{}
	}

	n = 0
	for _, item := range feed.Items {
		guid, err := getGUID(blog.ID, item.GUID)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}

		_, ok := exists[guid]
		if ok {
			continue
		}

		e := &record.Entry{
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Link:        item.Link,
			AuthorID:    blog.UserID,
			GUID:        guid,
			BlogID:      blog.ID,
		}
		if i := item.Image; i != nil {
			e.ImageURL = i.URL
		}
		if t := item.PublishedParsed; t == nil {
			e.PublishedAt = time.Now().In(boil.GetLocation())
		} else {
			e.PublishedAt = t.In(boil.GetLocation())
		}

		err = e.Insert(s.ctx, tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return 0, errors.WithStack(err)
		}
		n++
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return 0, errors.WithStack(err)
	}
	return n, nil
}

func getGUID(blogID int64, guid string) (string, error) {
	h := sha256.New()
	err := binary.Write(h, binary.LittleEndian, blogID)
	if err != nil {
		return "", errors.WithStack(err)
	}
	_, err = h.Write([]byte(guid))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return base64.RawURLEncoding.EncodeToString(h.Sum(nil)), nil
}
