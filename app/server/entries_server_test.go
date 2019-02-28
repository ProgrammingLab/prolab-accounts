package server

import (
	"context"
	"log"
	"testing"

	"github.com/ProgrammingLab/prolab-accounts/model"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

func Test_EntriesServer_ListPublicEntries(t *testing.T) {
	cfg, err := config.LoadConfig("../../.env")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	store := di.MustCreateTestStoreComponent(cfg)
	defer store.MustClose()
	ctx := context.Background()
	users, err := CreateTestUsers(ctx, store)
	if err != nil {
		t.Fatal(err)
	}
	blogs, err := CreateTestBlogs(ctx, store, users)
	if err != nil {
		t.Fatal(err)
	}
	err = CreateTestEntries(ctx, store, blogs)
	if err != nil {
		t.Fatal(err)
	}

	svr := NewEntryServiceServer(store, cfg)
	req := &api_pb.ListEntriesRequest{}

	resp, err := svr.ListPublicEntries(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}

	if got, want := resp.NextPageToken, 0; got != int64(want) {
		t.Errorf("NextPageToken is %v, want %v", got, want)
	}

	for _, e := range resp.Entries {
		if got, want := e.Author.Email, ""; got != want {
			t.Errorf("Author.Email is %v, want %v", got, want)
		}
		if got, want := e.Author.FullName, ""; got != want {
			t.Errorf("Author.FullName is %v, want %v", got, want)
		}
		if got, want := e.Author.Description, ""; got != want {
			t.Errorf("Author.Description is %v, want %v", got, want)
		}
		if got, want := userProfileScope(model.UserID(e.Author.UserId)), model.Public; got != want {
			t.Errorf("Author's profile scope is %v, want %v", got, want)
		}
	}
}
