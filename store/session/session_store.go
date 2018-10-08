package sessionstore

import (
	"github.com/ProgrammingLab/prolab-accounts/store"
	"github.com/go-redis/redis"
)

type sessionStoreImpl struct {
	client *redis.Client
}

// NewSessionStore returns new session store
func NewSessionStore(cli *redis.Client) store.SessionStore {
	return &sessionStoreImpl{client: cli}
}
