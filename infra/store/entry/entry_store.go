package entrystore

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/binary"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
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

func (s *entryStoreImpl) CreateEntries(blog *record.Blog, feed *gofeed.Feed) (n int64, err error) {
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
		if t := item.PublishedParsed; t != nil {
			e.PublishedAt = null.TimeFrom(t.In(boil.GetLocation()))
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
