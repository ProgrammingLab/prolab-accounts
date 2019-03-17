package di

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq" // for database/sql
	minio "github.com/minio/minio-go"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	departmentstore "github.com/ProgrammingLab/prolab-accounts/infra/store/department"
	emailconfirmationstore "github.com/ProgrammingLab/prolab-accounts/infra/store/email_confirmation"
	entrystore "github.com/ProgrammingLab/prolab-accounts/infra/store/entry"
	feedstore "github.com/ProgrammingLab/prolab-accounts/infra/store/feed"
	githubstore "github.com/ProgrammingLab/prolab-accounts/infra/store/github"
	heartbeatstore "github.com/ProgrammingLab/prolab-accounts/infra/store/heartbeat"
	invitationstore "github.com/ProgrammingLab/prolab-accounts/infra/store/invitation"
	resetstore "github.com/ProgrammingLab/prolab-accounts/infra/store/password_reset"
	profilestore "github.com/ProgrammingLab/prolab-accounts/infra/store/profile"
	rolestore "github.com/ProgrammingLab/prolab-accounts/infra/store/role"
	sessionstore "github.com/ProgrammingLab/prolab-accounts/infra/store/session"
	userstore "github.com/ProgrammingLab/prolab-accounts/infra/store/user"
	userblogstore "github.com/ProgrammingLab/prolab-accounts/infra/store/user_blog"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

// StoreComponent is an interface of stores
type StoreComponent interface {
	UserStore(ctx context.Context) store.UserStore
	SessionStore(ctx context.Context) store.SessionStore
	ProfileStore(ctx context.Context) store.ProfileStore
	UserBlogStore(ctx context.Context) store.UserBlogStore
	FeedStore(ctx context.Context) store.FeedStore
	EntryStore(ctx context.Context) store.EntryStore
	RoleStore(ctx context.Context) store.RoleStore
	HeartbeatStore(ctx context.Context) store.HeartbeatStore
	DepartmentStore(ctx context.Context) store.DepartmentStore
	InvitationStore(ctx context.Context) store.InvitationStore
	GitHubStore(ctx context.Context) store.GitHubStore
	EmailConfirmationStore(ctx context.Context) store.EmailConfirmationStore
	PasswordResetStore(ctx context.Context) store.PasswordResetStore
}

// NewStoreComponent returns new store component
func NewStoreComponent(cfg *config.Config) (StoreComponent, error) {
	db, err := connectRDB(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := connectRedis(cfg)
	if err != nil {
		return nil, err
	}

	min, err := connectMinio(cfg)
	if err != nil {
		return nil, err
	}

	return &storeComponentImpl{
		db:       sqlutil.New(db),
		client:   cli,
		minioCli: min,
		cfg:      cfg,
	}, nil
}

func connectRDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DataBaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "faild to connect db")
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, errors.Wrap(err, "faild to ping db")
	}

	return db, nil
}

func connectRedis(cfg *config.Config) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
		DB:   0,
	})
	_, err := cli.Ping().Result()
	if err != nil {
		return nil, errors.Wrap(err, "faild to connect redis")
	}

	return cli, nil
}

func connectMinio(cfg *config.Config) (*minio.Client, error) {
	cli, err := minio.New(cfg.MinioEndpoint, cfg.MinioAccessKey, cfg.MinioSecretKey, false)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	name := cfg.MinioBucketName
	ex, err := cli.BucketExists(name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if ex {
		return cli, nil
	}

	err = cli.MakeBucket(name, "asia-northeast1")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	policy := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Action": ["s3:GetObject"],
				"Effect": "Allow",
				"Principal": {"AWS": ["*"]},
				"Resource": ["arn:aws:s3:::` + name + `/*"],
				"Sid": ""
			}
		]
	}`
	err = cli.SetBucketPolicy(name, policy)
	if err != nil {
		_ = cli.RemoveBucket(name)
		return nil, errors.WithStack(err)
	}

	return cli, nil
}

type storeComponentImpl struct {
	db       *sqlutil.DB
	client   *redis.Client
	minioCli *minio.Client
	cfg      *config.Config
}

func (s *storeComponentImpl) UserStore(ctx context.Context) store.UserStore {
	return userstore.NewUserStore(ctx, s.db, s.minioCli, s.cfg.MinioBucketName)
}

func (s *storeComponentImpl) SessionStore(ctx context.Context) store.SessionStore {
	return sessionstore.NewSessionStore(ctx, s.client, s.db)
}

func (s *storeComponentImpl) ProfileStore(ctx context.Context) store.ProfileStore {
	return profilestore.NewProfileStore(ctx, s.db)
}

func (s *storeComponentImpl) UserBlogStore(ctx context.Context) store.UserBlogStore {
	return userblogstore.NewUserBlogStore(ctx, s.db)
}

func (s *storeComponentImpl) FeedStore(ctx context.Context) store.FeedStore {
	return feedstore.NewFeedStore(ctx)
}

func (s *storeComponentImpl) EntryStore(ctx context.Context) store.EntryStore {
	return entrystore.NewEntryStore(ctx, s.db)
}

func (s *storeComponentImpl) RoleStore(ctx context.Context) store.RoleStore {
	return rolestore.NewRoleStore(ctx, s.db)
}

func (s *storeComponentImpl) DepartmentStore(ctx context.Context) store.DepartmentStore {
	return departmentstore.NewDepartmentStore(ctx, s.db)
}

func (s *storeComponentImpl) HeartbeatStore(ctx context.Context) store.HeartbeatStore {
	return heartbeatstore.NewHeartbeatStore(ctx, s.client, s.cfg)
}

func (s *storeComponentImpl) InvitationStore(ctx context.Context) store.InvitationStore {
	return invitationstore.NewInvitationStore(ctx, s.db)
}

func (s *storeComponentImpl) GitHubStore(ctx context.Context) store.GitHubStore {
	return githubstore.NewGitHubStore(ctx, s.db, s.client)
}

func (s *storeComponentImpl) EmailConfirmationStore(ctx context.Context) store.EmailConfirmationStore {
	return emailconfirmationstore.NewEmailConfirmationStore(ctx, s.db)
}

func (s *storeComponentImpl) PasswordResetStore(ctx context.Context) store.PasswordResetStore {
	return resetstore.NewPasswordResetStore(ctx, s.db)
}
