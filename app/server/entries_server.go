package server

import (
	"context"
	"database/sql"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// EntryServiceServer is a composite interface of api_pb.EntryServiceServer and grapiserver.Server.
type EntryServiceServer interface {
	api_pb.EntryServiceServer
	grapiserver.Server
}

// NewEntryServiceServer creates a new EntryServiceServer instance.
func NewEntryServiceServer(store di.StoreComponent, cfg *config.Config) EntryServiceServer {
	return &entryServiceServerImpl{
		StoreComponent: store,
		cfg:            cfg,
	}
}

type entryServiceServerImpl struct {
	di.StoreComponent
	cfg *config.Config
}

func (s *entryServiceServerImpl) ListPublicEntries(ctx context.Context, req *api_pb.ListEntriesRequest) (*api_pb.ListEntriesResponse, error) {
	size := req.GetPageSize()
	if size == 0 {
		size = 50
	}
	nano := req.GetPageToken()
	var t time.Time
	if nano == 0 {
		t = time.Now()
	} else {
		t = time.Unix(0, nano)
	}

	es := s.EntryStore(ctx)
	entries, next, err := es.ListPublicEntries(t, int(size))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return &api_pb.ListEntriesResponse{}, nil
		}
		return nil, err
	}

	resp := entriesToResponse(entries, false, s.cfg)
	return &api_pb.ListEntriesResponse{
		Entries:       resp,
		NextPageToken: next.UnixNano(),
	}, nil
}

func entriesToResponse(entries []*record.Entry, includePrivate bool, cfg *config.Config) []*api_pb.Entry {
	res := make([]*api_pb.Entry, 0, len(entries))
	for _, e := range entries {
		res = append(res, entryToResponse(e, includePrivate, cfg))
	}

	return res
}

func entryToResponse(entry *record.Entry, includePrivate bool, cfg *config.Config) *api_pb.Entry {
	e := &api_pb.Entry{
		EntryId:     uint32(entry.ID),
		Title:       entry.Title,
		Description: entry.Description,
		Content:     entry.Content,
		Link:        entry.Link,
		ImageUrl:    entry.ImageURL,
		UpdatedAt:   timeToResponse(entry.UpdatedAt),
		PublishedAt: timeToResponse(entry.PublishedAt),
	}
	if r := entry.R; r != nil {
		e.Author = userToResponse(r.Author, includePrivate, cfg)
		e.Blog = blogToResponse(r.Blog)
	}

	return e
}

func timeToResponse(t time.Time) *types.Timestamp {
	t = t.UTC()
	return &types.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
