package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// DepartmentServiceServer is a composite interface of api_pb.DepartmentServiceServer and grapiserver.Server.
type DepartmentServiceServer interface {
	api_pb.DepartmentServiceServer
	grapiserver.Server
}

// NewDepartmentServiceServer creates a new DepartmentServiceServer instance.
func NewDepartmentServiceServer() DepartmentServiceServer {
	return &departmentServiceServerImpl{}
}

type departmentServiceServerImpl struct {
}

func (s *departmentServiceServerImpl) ListDepartments(ctx context.Context, req *api_pb.ListDepartmentsRequest) (*api_pb.ListDepartmentsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *departmentServiceServerImpl) GetDepartment(ctx context.Context, req *api_pb.GetDepartmentRequest) (*api_pb.Department, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
