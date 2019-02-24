package feedstore

import (
	"net/http"
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

type feedURLResult struct {
	FeedURL string
	Error   error
}

func getFeedURLWithSuffixes(url string, cli *http.Client) (string, error) {
	c := make(chan *feedURLResult)
	for _, s := range suffixes {
		go func(suffix string, cli *http.Client) {
			u, err := getFeedURL(url, suffix, cli)
			if err == nil {
				c <- &feedURLResult{
					FeedURL: u,
					Error:   nil,
				}
				return
			}

			c <- &feedURLResult{
				FeedURL: "",
				Error:   err,
			}
		}(s, cli)
	}

	var feed string
	err := ErrFeedURLNotFound
	for range suffixes {
		res := <-c
		if err == nil || res.Error != nil {
			continue
		}

		feed = res.FeedURL
		err = res.Error
	}

	return feed, err
}

const (
	mediumPrefix = "https://medium.com/@"
)

func getMediumFeed(url string, cli *http.Client) (string, error) {
	if !strings.HasPrefix(url, mediumPrefix) {
		return "", ErrFeedURLNotFound
	}

	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	name := url[len(mediumPrefix):]
	feed := "https://medium.com/feed/@" + name

	p := gofeed.NewParser()
	p.Client = cli
	_, err := p.ParseURL(feed)
	if err != nil {
		return "", err
	}
	return feed, nil
}

func getFeedURL(url, suffix string, cli *http.Client) (string, error) {
	if url[len(url)-1] != '/' {
		url += "/"
	}

	url += suffix
	p := gofeed.NewParser()
	p.Client = cli
	_, err := p.ParseURL(url)
	if err != nil {
		return "", err
	}

	return url, nil
}
