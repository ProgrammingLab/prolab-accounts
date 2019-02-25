package server

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
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

func (s *entryServiceServerImpl) ListEntries(ctx context.Context, req *api_pb.ListEntriesRequest) (*api_pb.ListEntriesResponse, error) {
	es := s.EntryStore(ctx)
	entries, err := es.ListPublicEntries()
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return &api_pb.ListEntriesResponse{}, nil
		}
		return nil, err
	}

	resp := entriesToResponse(entries, false, s.cfg)
	return &api_pb.ListEntriesResponse{
		Entries: resp,
	}, nil
}

func entriesToResponse(entries []*record.Entry, includeEmail bool, cfg *config.Config) []*api_pb.Entry {
	res := make([]*api_pb.Entry, 0, len(entries))
	for _, e := range entries {
		res = append(res, entryToResponse(e, includeEmail, cfg))
	}

	return res
}

func entryToResponse(entry *record.Entry, includeEmail bool, cfg *config.Config) *api_pb.Entry {
	e := &api_pb.Entry{
		EntryId:     uint32(entry.ID),
		Title:       entry.Title,
		Description: entry.Description,
		Content:     entry.Content,
		Link:        entry.Link,
		ImageUrl:    entry.ImageURL,
		UpdatedAt:   timeToResponse(entry.UpdatedAt),
	}
	if t := entry.PublishedAt; t.Valid {
		e.PublishedAt = timeToResponse(t.Time)
	}
	if r := entry.R; r != nil {
		e.Author = userToResponse(r.Author, includeEmail, cfg)
		e.Blog = blogToResponse(r.Blog)
	}

	return e
}

func timeToResponse(t time.Time) *timestamp.Timestamp {
	t = t.UTC()
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
