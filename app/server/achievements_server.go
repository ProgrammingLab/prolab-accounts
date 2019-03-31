package server

import (
	"context"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
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

const (
	pageTokenTimeLayout = time.RFC3339
)

var (
	// ErrInvalidPageToken is returned when page token is invalid
	ErrInvalidPageToken = status.Error(codes.InvalidArgument, "invalid page token")
)

func (s *achievementServiceServerImpl) ListAchievements(ctx context.Context, req *api_pb.ListAchievementsRequest) (*api_pb.ListAchievementsResponse, error) {
	_, ok := interceptor.GetCurrentUserID(ctx)

	size := req.GetPageSize()
	if size == 0 {
		size = 50
	}

	var (
		before time.Time
		err    error
	)
	if token := req.GetPageToken(); token == "" {
		before = time.Now()
	} else {
		before, err = time.Parse(pageTokenTimeLayout, req.GetPageToken())
		if err != nil {
			return nil, ErrInvalidPageToken
		}
	}

	as := s.AchievementStore(ctx)
	aches, next, err := as.ListAchievements(before, int(size))
	if err != nil {
		return nil, err
	}

	return &api_pb.ListAchievementsResponse{
		Achievements:  achievementsToResponse(aches, ok, s.cfg),
		NextPageToken: next.Format(pageTokenTimeLayout),
	}, nil
}

func (s *achievementServiceServerImpl) GetAchievement(context.Context, *api_pb.GetAchievementRequest) (*api_pb.Achievement, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *achievementServiceServerImpl) CreateAchievement(ctx context.Context, req *api_pb.CreateAchievementRequest) (*api_pb.Achievement, error) {
	_, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	ach := req.GetAchievement()
	rec := &record.Achievement{
		Title:       ach.GetTitle(),
		Award:       ach.GetAward(),
		URL:         ach.GetUrl(),
		Description: ach.GetDescription(),
		HappenedAt:  toTime(ach.HappenedAt),
	}
	as := s.AchievementStore(ctx)
	err := as.CreateAchievement(rec, toUserIDs(ach.GetMembers()))
	if err != nil {
		return nil, err
	}

	return achievementToResponse(rec, true, s.cfg), nil
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

func toUserIDs(users []*api_pb.User) []model.UserID {
	res := make([]model.UserID, 0, len(users))
	for _, u := range users {
		res = append(res, model.UserID(u.GetUserId()))
	}
	return res
}

func toTime(t *types.Timestamp) time.Time {
	return time.Unix(t.GetSeconds(), int64(t.GetNanos()))
}

func achievementToResponse(ach *record.Achievement, includePrivate bool, cfg *config.Config) *api_pb.Achievement {
	resp := &api_pb.Achievement{
		AchievementId: uint32(ach.ID),
		Title:         ach.Title,
		Award:         ach.Award,
		Url:           ach.URL,
		Description:   ach.Description,
		ImageUrl:      "todo", // TODO
		HappenedAt:    timeToResponse(ach.HappenedAt),
	}

	members := make([]*record.User, 0)
	hidden := 0
	if ach.R != nil {
		hidden = len(ach.R.AchievementUsers)
		for _, au := range ach.R.AchievementUsers {
			if au.R == nil || au.R.User == nil {
				continue
			}

			u := au.R.User
			if !includePrivate && (u.R == nil || u.R.Profile == nil || u.R.Profile.ProfileScope.Int != int(model.Public)) {
				continue
			}
			members = append(members, u)
		}
		hidden -= len(members)
	}

	resp.Members = usersToResponse(members, false, cfg)
	resp.HiddenMembersCount = int32(hidden)

	return resp
}

func achievementsToResponse(aches []*record.Achievement, includePrivate bool, cfg *config.Config) []*api_pb.Achievement {
	resp := make([]*api_pb.Achievement, 0, len(aches))
	for _, a := range aches {
		resp = append(resp, achievementToResponse(a, includePrivate, cfg))
	}
	return resp
}
