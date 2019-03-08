package job

import (
	"context"
	"time"

	"google.golang.org/grpc/grpclog"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

func feedJob(ctx context.Context, store di.StoreComponent, cfg *config.Config) error {
	bs := store.UserBlogStore(ctx)
	blogs, err := bs.ListUserBlogs()
	if err != nil {
		return err
	}

	fs := store.FeedStore(ctx)
	es := store.EntryStore(ctx)
	for _, b := range blogs {
		if b.FeedURL == "" {
			continue
		}
		feed, err := fs.GetFeed(b.FeedURL)
		if err != nil {
			grpclog.Errorf("feed job: failed to get feed: blog id: %v : %+v", b.ID, err)
			continue
		}

		n, err := es.CreateEntries(b, feed)
		if err != nil {
			return err
		}
		if cfg.DebugLog {
			grpclog.Infof("feed job: created %v entries", n)
		}

		<-time.After(100 * time.Millisecond)
	}

	return nil
}
