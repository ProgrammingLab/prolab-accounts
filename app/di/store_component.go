package di

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq" // for database/sql
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

	return &storeComponentImpl{
		db:     db,
		client: cli,
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

type storeComponentImpl struct {
	db     *sql.DB
	client *redis.Client
}

func (s *storeComponentImpl) UserStore(ctx context.Context) store.UserStore {
	return userstore.NewUserStore(ctx, s.db)
}

func (s *storeComponentImpl) SessionStore(ctx context.Context) store.SessionStore {
	return sessionstore.NewSessionStore(ctx, s.client, s.db)
}

func (s *storeComponentImpl) ProfileStore(ctx context.Context) store.ProfileStore {
	return profilestore.NewProfileStore(ctx, s.db)
}
