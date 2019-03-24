package server

import (
	"context"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
)

// AchievementServiceServer is a composite interface of api_pb.AchievementServiceServer and grapiserver.Server.
type AchievementServiceServer interface {
	api_pb.AchievementServiceServer
	grapiserver.Server
}

// NewAchievementServiceServer creates a new AchievementServiceServer instance.
func NewAchievementServiceServer(store di.StoreComponent, cli di.ClientComponent, cfg *config.Config) AchievementServiceServer {
	return &achievementServiceServerImpl{
		StoreComponent:  store,
		ClientComponent: cli,
		cfg:             cfg,
	}
}

type achievementServiceServerImpl struct {
	di.StoreComponent
	di.ClientComponent
	cfg *config.Config
}

func (s *achievementServiceServerImpl) ListAchievements(context.Context, *api_pb.ListAchievementsRequest) (*api_pb.ListAchievementsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *achievementServiceServerImpl) GetAchievement(context.Context, *api_pb.GetAchievementRequest) (*api_pb.Achievement, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *achievementServiceServerImpl) CreateAchievement(context.Context, *api_pb.CreateAchievementRequest) (*api_pb.Achievement, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *achievementServiceServerImpl) UpdateAchievement(context.Context, *api_pb.UpdateAchievementRequest) (*api_pb.Achievement, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *achievementServiceServerImpl) UpdateAchievementImage(context.Context, *api_pb.UpdateAchievementImageRequest) (*api_pb.Achievement, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *achievementServiceServerImpl) DeleteAchievement(context.Context, *api_pb.DeleteAchievementRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
