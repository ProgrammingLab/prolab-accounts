package server

import (
	"context"

	"github.com/mmcdole/gofeed"
	"github.com/volatiletech/null"

	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

func CreateTestUsers(ctx context.Context, s di.StoreComponent) ([]*record.User, error) {
	users := []*record.User{
		{
			Name:     "hoge",
			Email:    "hoge@kurume-nct.com",
			FullName: "Hoge Hoge",
		},
		{
			Name:     "piyo",
			Email:    "piyo@kurume-nct.com",
			FullName: "Poyo Piyo",
		},
	}

	us := s.UserStore(ctx)
	for _, u := range users {
		err := us.CreateUser(u)
		if err != nil {
			return nil, err
		}
	}

	ps := s.ProfileStore(ctx)
	for _, u := range users {
		scope := userProfileScope(model.UserID(u.ID))
		p := &record.Profile{
			Description:  "hoge",
			ProfileScope: null.IntFrom(int(scope)),
		}

		err := ps.CreateOrUpdateProfile(model.UserID(u.ID), p, true)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func userProfileScope(userID model.UserID) model.ProfileScope {
	if userID%2 == 0 {
		return model.MembersOnly
	}
	return model.Public
}

func CreateTestBlogs(ctx context.Context, s di.StoreComponent, users []*record.User) ([]*record.Blog, error) {
	bs := s.UserBlogStore(ctx)

	blogs := make([]*record.Blog, 0, len(users))
	for _, u := range users {
		b := &record.Blog{
			URL:     "https://hogeblog.com/" + u.Name,
			FeedURL: "https://hogeblog.com/" + u.Name + "/feed",
			UserID:  u.ID,
		}
		err := bs.CreateUserBlog(b)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}

	return blogs, nil
}

func CreateTestEntries(ctx context.Context, s di.StoreComponent, blogs []*record.Blog) error {
	entries := []*gofeed.Item{
		{
			Title:       "title",
			Description: "hoge",
			Content:     "hogehoge",
		},
	}
	feed := &gofeed.Feed{
		Items: entries,
	}

	es := s.EntryStore(ctx)
	for _, b := range blogs {
		_, err := es.CreateEntries(b, feed)
		if err != nil {
			return err
		}
	}

	return nil
}
