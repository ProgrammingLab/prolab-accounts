package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

func Test_SessionServiceServer_ListSessions(t *testing.T) {
	svr := NewSessionServiceServer()

	ctx := context.Background()
	req := &api_pb.ListSessionsRequest{}

	resp, err := svr.ListSessions(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_SessionServiceServer_GetSession(t *testing.T) {
	svr := NewSessionServiceServer()

	ctx := context.Background()
	req := &api_pb.GetSessionRequest{}

	resp, err := svr.GetSession(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_SessionServiceServer_CreateSession(t *testing.T) {
	svr := NewSessionServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateSessionRequest{}

	resp, err := svr.CreateSession(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_SessionServiceServer_UpdateSession(t *testing.T) {
	svr := NewSessionServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateSessionRequest{}

	resp, err := svr.UpdateSession(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_SessionServiceServer_DeleteSession(t *testing.T) {
	svr := NewSessionServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteSessionRequest{}

	resp, err := svr.DeleteSession(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}
