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
	profilestore "github.com/ProgrammingLab/prolab-accounts/infra/store/profile"
	sessionstore "github.com/ProgrammingLab/prolab-accounts/infra/store/session"
	userstore "github.com/ProgrammingLab/prolab-accounts/infra/store/user"
)

// StoreComponent is an interface of stores
type StoreComponent interface {
	UserStore(ctx context.Context) store.UserStore
	SessionStore(ctx context.Context) store.SessionStore
	ProfileStore(ctx context.Context) store.ProfileStore
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
		db:       db,
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
	db       *sql.DB
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
