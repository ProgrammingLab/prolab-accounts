package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

func Test_EntryServiceServer_ListEntries(t *testing.T) {
	svr := NewEntryServiceServer()

	ctx := context.Background()
	req := &api_pb.ListEntriesRequest{}

	resp, err := svr.ListEntries(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_EntryServiceServer_GetEntry(t *testing.T) {
	svr := NewEntryServiceServer()

	ctx := context.Background()
	req := &api_pb.GetEntryRequest{}

	resp, err := svr.GetEntry(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_EntryServiceServer_CreateEntry(t *testing.T) {
	svr := NewEntryServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateEntryRequest{}

	resp, err := svr.CreateEntry(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_EntryServiceServer_UpdateEntry(t *testing.T) {
	svr := NewEntryServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateEntryRequest{}

	resp, err := svr.UpdateEntry(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_EntryServiceServer_DeleteEntry(t *testing.T) {
	svr := NewEntryServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteEntryRequest{}

	resp, err := svr.DeleteEntry(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
