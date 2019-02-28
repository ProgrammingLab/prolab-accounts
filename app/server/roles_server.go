package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// RoleServiceServer is a composite interface of api_pb.RoleServiceServer and grapiserver.Server.
type RoleServiceServer interface {
	api_pb.RoleServiceServer
	grapiserver.Server
}

// NewRoleServiceServer creates a new RoleServiceServer instance.
func NewRoleServiceServer() RoleServiceServer {
	return &roleServiceServerImpl{}
}

type roleServiceServerImpl struct {
}

func (s *roleServiceServerImpl) ListRoles(ctx context.Context, req *api_pb.ListRolesRequest) (*api_pb.ListRolesResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *roleServiceServerImpl) GetRole(ctx context.Context, req *api_pb.GetRoleRequest) (*api_pb.Role, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *roleServiceServerImpl) CreateRole(ctx context.Context, req *api_pb.CreateRoleRequest) (*api_pb.Role, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *roleServiceServerImpl) UpdateRole(ctx context.Context, req *api_pb.UpdateRoleRequest) (*api_pb.Role, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *roleServiceServerImpl) DeleteRole(ctx context.Context, req *api_pb.DeleteRoleRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
