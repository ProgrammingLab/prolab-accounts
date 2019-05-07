package server

import (
	"context"
	"database/sql"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
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
	// MaxImageSize represents max of image size
	MaxImageSize        = 1024 * 1024 * 3
	pageTokenTimeLayout = time.RFC3339
)

var (
	// ErrInvalidPageToken is returned when page token is invalid
	ErrInvalidPageToken = status.Error(codes.InvalidArgument, "invalid page token")
	// ErrImageSizeTooLarge will be returned when the image is too large
	ErrImageSizeTooLarge = status.Error(codes.InvalidArgument, "image must be smaller than 3MiB")
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

func (s *achievementServiceServerImpl) GetAchievement(ctx context.Context, req *api_pb.GetAchievementRequest) (*api_pb.Achievement, error) {
	_, ok := interceptor.GetCurrentUserID(ctx)
	as := s.AchievementStore(ctx)
	ach, err := as.GetAchievement(int64(req.GetAchievementId()))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return achievementToResponse(ach, ok, s.cfg), nil
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
	rec, err := as.CreateAchievement(rec, toUserIDs(ach.GetMembers()))
	if err != nil {
		return nil, err
	}

	return achievementToResponse(rec, true, s.cfg), nil
}

func (s *achievementServiceServerImpl) UpdateAchievement(ctx context.Context, req *api_pb.UpdateAchievementRequest) (*api_pb.Achievement, error) {
	_, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	ach := req.GetAchievement()

	// ここの間かな
	if ach.ImageUrl == "" {
		// 非同期処理
		go func() {
			resp, err := http.Get(ach.Url)
			if err != nil {
				grpclog.Errorf("ogImage: %+v", err)
				return
			}

			defer func() { _ = resp.Body.Close() }()

			og := opengraph.NewOpenGraph()
			err = og.ProcessHTML(resp.Body)
			if err != nil {
				grpclog.Errorf("ogImage: %+v", err)
				return
			}

			images := og.Images
			if len(images) == 0 {
				return
			}

			image := images[0]
			imgResp, err := http.Get(image.SecureURL)
			if err != nil {
				grpclog.Errorf("ogImage: %+v", err)
				return
			}

			imageByte, err := ioutil.ReadAll(imgResp.Body)
			if err != nil {
				grpclog.Errorf("ogImage: %+v", err)
				return
			}

			is := s.ImageStore(context.Background())
			name, err := is.CreateImage(imageByte)
			if err != nil {
				grpclog.Errorf("ogImage: %+v", err)
				return
			}

			_, _, err = s.AchievementStore(context.Background()).UpdateAchievementImage(int64(req.Achievement.AchievementId), name)
			if err != nil {
				grpclog.Errorf("ogImage: %+v", err)
				return
			}
		}()
	}

	rec := &record.Achievement{
		ID:          int64(ach.AchievementId),
		Title:       ach.GetTitle(),
		Award:       ach.GetAward(),
		URL:         ach.GetUrl(),
		Description: ach.GetDescription(),
		HappenedAt:  toTime(ach.HappenedAt),
	}
	as := s.AchievementStore(ctx)
	rec, err := as.UpdateAchievement(rec, toUserIDs(ach.GetMembers()))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return achievementToResponse(rec, true, s.cfg), nil
}

func (s *achievementServiceServerImpl) UpdateAchievementImage(ctx context.Context, req *api_pb.UpdateAchievementImageRequest) (*api_pb.Achievement, error) {
	_, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	image := req.GetImage()
	if MaxImageSize < len(image) {
		return nil, ErrImageSizeTooLarge
	}

	is := s.ImageStore(ctx)
	name, err := is.CreateImage(image)
	if err != nil {
		return nil, err
	}

	as := s.AchievementStore(ctx)
	ach, old, err := as.UpdateAchievementImage(int64(req.GetAchievementId()), name)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	go func() {
		if old == "" {
			return
		}

		is := s.ImageStore(context.Background())
		err := is.DeleteImage(old)
		if err != nil {
			grpclog.Errorf("failed to delete old user icon: %+v", err)
		}
	}()

	return achievementToResponse(ach, true, s.cfg), nil
}

func (s *achievementServiceServerImpl) DeleteAchievement(ctx context.Context, req *api_pb.DeleteAchievementRequest) (*empty.Empty, error) {
	_, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	as := s.AchievementStore(ctx)
	err := as.DeleteAchievement(int64(req.GetAchievementId()))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return &empty.Empty{}, nil
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
		HappenedAt:    timeToResponse(ach.HappenedAt),
	}

	if ach.ImageFilename.Valid {
		resp.ImageUrl = cfg.MinioPublicURL + "/" + cfg.MinioBucketName + "/" + ach.ImageFilename.String
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
