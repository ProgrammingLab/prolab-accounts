package feedstore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
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

type feedURLGetter func(blogURL string, cli *http.Client) (feed string, err error)

var (
	// ErrFeedURLNotFound will be returned when feed url not found
	ErrFeedURLNotFound = fmt.Errorf("feed url not found")

	feedURLGetters = []feedURLGetter{
		getMediumFeed,
		getFeedURLWithSuffixes,
	}
)

func (s *feedStoreImpl) GetFeedURL(url string) (string, error) {
	for _, g := range feedURLGetters {
		u, err := g(url, &http.Client{})
		if err == nil {
			return u, nil
		}
	}

	return "", ErrFeedURLNotFound
}

func (s *feedStoreImpl) IsValidFeedURL(feedURL string) error {
	_, err := s.GetFeed(feedURL)
	return err
}

func (s *feedStoreImpl) GetFeed(feedURL string) (*gofeed.Feed, error) {
	p := gofeed.NewParser()
	f, err := p.ParseURL(feedURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return f, nil
}
