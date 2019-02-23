package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// EntryServiceServer is a composite interface of api_pb.EntryServiceServer and grapiserver.Server.
type EntryServiceServer interface {
	api_pb.EntryServiceServer
	grapiserver.Server
}

// NewEntryServiceServer creates a new EntryServiceServer instance.
func NewEntryServiceServer() EntryServiceServer {
	return &entryServiceServerImpl{}
}

type entryServiceServerImpl struct {
}

func (s *entryServiceServerImpl) ListEntries(ctx context.Context, req *api_pb.ListEntriesRequest) (*api_pb.ListEntriesResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
