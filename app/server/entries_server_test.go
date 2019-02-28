package server

import (
	"context"
	"log"
	"testing"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

func Test_EntriesServer_ListPublicEntries(t *testing.T) {
	cfg, err := config.LoadConfig("../../.env")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	store := di.MustCreateTestStoreComponent(cfg)
	defer store.MustClose()
	svr := NewEntryServiceServer(store, cfg)

	ctx := context.Background()
	req := &api_pb.ListEntriesRequest{}

	_, err = svr.ListPublicEntries(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}
}
