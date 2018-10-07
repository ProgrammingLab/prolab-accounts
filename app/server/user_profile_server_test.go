package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

func Test_UserProfileServiceServer_ListUserProfiles(t *testing.T) {
	svr := NewUserProfileServiceServer()

	ctx := context.Background()
	req := &api_pb.ListUserProfilesRequest{}

	resp, err := svr.ListUserProfiles(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_UserProfileServiceServer_GetUserProfile(t *testing.T) {
	svr := NewUserProfileServiceServer()

	ctx := context.Background()
	req := &api_pb.GetUserProfileRequest{}

	resp, err := svr.GetUserProfile(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_UserProfileServiceServer_CreateUserProfile(t *testing.T) {
	svr := NewUserProfileServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateUserProfileRequest{}

	resp, err := svr.CreateUserProfile(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_UserProfileServiceServer_UpdateUserProfile(t *testing.T) {
	svr := NewUserProfileServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateUserProfileRequest{}

	resp, err := svr.UpdateUserProfile(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_UserProfileServiceServer_DeleteUserProfile(t *testing.T) {
	svr := NewUserProfileServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteUserProfileRequest{}

	resp, err := svr.DeleteUserProfile(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}
