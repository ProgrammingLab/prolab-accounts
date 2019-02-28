package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

func Test_RoleServiceServer_ListRoles(t *testing.T) {
	svr := NewRoleServiceServer()

	ctx := context.Background()
	req := &api_pb.ListRolesRequest{}

	resp, err := svr.ListRoles(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_RoleServiceServer_GetRole(t *testing.T) {
	svr := NewRoleServiceServer()

	ctx := context.Background()
	req := &api_pb.GetRoleRequest{}

	resp, err := svr.GetRole(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_RoleServiceServer_CreateRole(t *testing.T) {
	svr := NewRoleServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateRoleRequest{}

	resp, err := svr.CreateRole(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_RoleServiceServer_UpdateRole(t *testing.T) {
	svr := NewRoleServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateRoleRequest{}

	resp, err := svr.UpdateRole(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_RoleServiceServer_DeleteRole(t *testing.T) {
	svr := NewRoleServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteRoleRequest{}

	resp, err := svr.DeleteRole(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
