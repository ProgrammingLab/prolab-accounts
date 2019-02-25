package store

import (
	"time"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/mmcdole/gofeed"
)

// EntryStore accesses entry data
type EntryStore interface {
	ListPublicEntries(before time.Time, limit int) ([]*record.Entry, time.Time, error)
	CreateEntries(blog *record.Blog, feed *gofeed.Feed) (int64, error)
}
