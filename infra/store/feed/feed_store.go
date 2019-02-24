package feedstore

import (
	"context"
	"fmt"

	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/mmcdole/gofeed"
)

type feedStoreImpl struct {
	ctx context.Context
}

// NewFeedStore returns new feed store
func NewFeedStore(ctx context.Context) store.FeedStore {
	return &feedStoreImpl{
		ctx: ctx,
	}
}

type feedURLGetter func(blogURL string) (feed string, err error)

var (
	// ErrFeedURLNotFound will be returned when feed url not found
	ErrFeedURLNotFound = fmt.Errorf("feed url not found")

	feedURLGetters = []feedURLGetter{
		getFeedURLWithSuffixes,
		getMediumFeed,
	}
)

func (s *feedStoreImpl) GetFeedURL(url string) (string, error) {
	for _, g := range feedURLGetters {
		u, err := g(url)
		if err == nil {
			return u, nil
		}
	}

	return "", ErrFeedURLNotFound
}

func (s *feedStoreImpl) GetFeed(feedURL string) (*gofeed.Feed, error) {
	panic("not implemented")
}
