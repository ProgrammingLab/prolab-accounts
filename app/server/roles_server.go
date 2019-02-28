package server

import (
	"context"
	"database/sql"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// RoleServiceServer is a composite interface of api_pb.RoleServiceServer and grapiserver.Server.
type RoleServiceServer interface {
	api_pb.RoleServiceServer
	grapiserver.Server
}

// NewRoleServiceServer creates a new RoleServiceServer instance.
func NewRoleServiceServer(store di.StoreComponent) RoleServiceServer {
	return &roleServiceServerImpl{
		StoreComponent: store,
	}
}

type roleServiceServerImpl struct {
	di.StoreComponent
}

func (s *roleServiceServerImpl) ListRoles(ctx context.Context, req *api_pb.ListRolesRequest) (*api_pb.ListRolesResponse, error) {
	rs := s.RoleStore(ctx)
	roles, err := rs.ListRoles()
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return &api_pb.ListRolesResponse{
				Roles: rolesToResponse(nil),
			}, nil
		}

		return nil, err
	}

	return &api_pb.ListRolesResponse{
		Roles: rolesToResponse(roles),
	}, nil
}

func (s *roleServiceServerImpl) GetRole(ctx context.Context, req *api_pb.GetRoleRequest) (*api_pb.Role, error) {
	rs := s.RoleStore(ctx)
	r, err := rs.GetRole(int64(req.GetRoleId()))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}

		return nil, err
	}

	return roleToResponse(r), nil
}

func rolesToResponse(roles []*record.Role) []*api_pb.Role {
	resp := make([]*api_pb.Role, 0, len(roles))
	for _, r := range roles {
		resp = append(resp, roleToResponse(r))
	}

	return resp
}

func roleToResponse(role *record.Role) *api_pb.Role {
	return &api_pb.Role{
		RoleId: uint32(role.ID),
		Name:   role.Name.String,
	}
}
