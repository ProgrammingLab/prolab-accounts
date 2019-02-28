package server

import (
	"context"
	"database/sql"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// DepartmentServiceServer is a composite interface of api_pb.DepartmentServiceServer and grapiserver.Server.
type DepartmentServiceServer interface {
	api_pb.DepartmentServiceServer
	grapiserver.Server
}

// NewDepartmentServiceServer creates a new DepartmentServiceServer instance.
func NewDepartmentServiceServer(store di.StoreComponent) DepartmentServiceServer {
	return &departmentServiceServerImpl{
		StoreComponent: store,
	}
}

type departmentServiceServerImpl struct {
	di.StoreComponent
}

func (s *departmentServiceServerImpl) ListDepartments(ctx context.Context, req *api_pb.ListDepartmentsRequest) (*api_pb.ListDepartmentsResponse, error) {
	ds := s.DepartmentStore(ctx)
	deps, err := ds.ListDepartments()
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return &api_pb.ListDepartmentsResponse{
				Departments: departmentsToResponse(nil),
			}, nil
		}
		return nil, err
	}

	return &api_pb.ListDepartmentsResponse{
		Departments: departmentsToResponse(deps),
	}, nil
}

func (s *departmentServiceServerImpl) GetDepartment(ctx context.Context, req *api_pb.GetDepartmentRequest) (*api_pb.Department, error) {
	ds := s.DepartmentStore(ctx)
	dep, err := ds.GetDepartment(int64(req.GetDepartmentId()))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return departmentToResponse(dep), nil
}

func departmentsToResponse(deps []*record.Department) []*api_pb.Department {
	resp := make([]*api_pb.Department, 0, len(deps))
	for _, d := range deps {
		resp = append(resp, departmentToResponse(d))
	}
	return resp
}

func departmentToResponse(dep *record.Department) *api_pb.Department {
	return &api_pb.Department{
		DepartmentId: uint32(dep.ID),
		Name:         dep.Name,
		ShortName:    dep.ShortName,
	}
}
