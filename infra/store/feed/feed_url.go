package feedstore

import (
	"strings"

	"github.com/mmcdole/gofeed"
)

var suffixes = []string{
	// Qiita
	"feed.atom",
	// note.mu
	"atom",
	// Hatena blog
	"feed",
	// Word Press
	"?feed=atom",
	// Excite blog
	"index.xml",
}

func getFeedURLWithSuffixes(url string) (string, error) {
	for _, s := range suffixes {
		u, err := getFeedURL(url, s)
		if err == nil {
			return u, nil
		}
	}

	return "", ErrFeedURLNotFound
}

const (
	mediumPrefix = "https://medium.com/@"
)

func getMediumFeed(url string) (string, error) {
	if !strings.HasPrefix(url, mediumPrefix) {
		return "", ErrFeedURLNotFound
	}

	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	name := url[len(mediumPrefix):]
	feed := "https://medium.com/feed/@" + name

	p := gofeed.NewParser()
	_, err := p.ParseURL(feed)
	if err != nil {
		return "", err
	}
	return feed, nil
}

func getFeedURL(url, suffix string) (string, error) {
	if url[len(url)-1] != '/' {
		url += "/"
	}

	url += suffix
	p := gofeed.NewParser()
	_, err := p.ParseURL(url)
	if err != nil {
		return "", err
	}

	return url, nil
}
