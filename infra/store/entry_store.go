package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/mmcdole/gofeed"
)

// EntryStore accesses entry data
type EntryStore interface {
	ListPublicEntries() ([]*record.Entry, error)
	CreateEntries(blog *record.Blog, feed *gofeed.Feed) (int64, error)
}
