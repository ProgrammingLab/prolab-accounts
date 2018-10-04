package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

func Test_UserServiceServer_ListUsers(t *testing.T) {
	svr := NewUserServiceServer()

	ctx := context.Background()
	req := &api_pb.ListUsersRequest{}

	resp, err := svr.ListUsers(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_UserServiceServer_GetUser(t *testing.T) {
	svr := NewUserServiceServer()

	ctx := context.Background()
	req := &api_pb.GetUserRequest{}

	resp, err := svr.GetUser(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_UserServiceServer_CreateUser(t *testing.T) {
	svr := NewUserServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateUserRequest{}

	resp, err := svr.CreateUser(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_UserServiceServer_UpdateUser(t *testing.T) {
	svr := NewUserServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateUserRequest{}

	resp, err := svr.UpdateUser(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_UserServiceServer_DeleteUser(t *testing.T) {
	svr := NewUserServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteUserRequest{}

	resp, err := svr.DeleteUser(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}
