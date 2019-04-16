package server

import (
	"context"
	"log"
	"testing"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

func Test_UsersServer_GetUser(t *testing.T) {
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

	t.Run("Get public user", func(t *testing.T) {
		var req *api_pb.GetUserRequest
		for _, u := range users {
			if userProfileScope(model.UserID(u.ID)) != model.Public {
				continue
			}

			req = &api_pb.GetUserRequest{
				UserName: u.Name,
			}
			break
		}

		srv := NewUserServiceServer(store, cfg)
		resp, err := srv.GetUser(ctx, req)

		if err != nil {
			t.Errorf("returned an error %v", err)
		}

		if resp == nil {
			t.Error("response should not nil")
		}

		if got, want := resp.Email, ""; got != want {
			t.Errorf("Email is %v, want %v", got, want)
		}
		if got, want := resp.FullName, ""; got != want {
			t.Errorf("FullName is %v, want %v", got, want)
		}
	})

	t.Run("Get members only user", func(t *testing.T) {
		var req *api_pb.GetUserRequest
		for _, u := range users {
			if userProfileScope(model.UserID(u.ID)) != model.MembersOnly {
				continue
			}

			req = &api_pb.GetUserRequest{
				UserName: u.Name,
			}
			break
		}

		srv := NewUserServiceServer(store, cfg)
		resp, err := srv.GetUser(ctx, req)

		if got, want := err, util.ErrNotFound; got != want {
			t.Errorf("error shold be %v, but %v", want, got)
		}

		if got := resp; got != nil {
			t.Errorf("resp shold be %v, but %v", nil, got)
		}
	})
}
