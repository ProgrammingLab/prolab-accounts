package sessionstore

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/store"
)

type sessionStoreImpl struct {
	ctx    context.Context
	client *redis.Client
}

const (
	// SessionExpiration is sessions expiration
	SessionExpiration = time.Hour * 24 * 7
)

// NewSessionStore returns new session store
func NewSessionStore(ctx context.Context, cli *redis.Client) store.SessionStore {
	return &sessionStoreImpl{
		ctx:    ctx,
		client: cli,
	}
}

func (s *sessionStoreImpl) CreateSession(userID model.UserID) (*model.Session, error) {
	session, err := model.NewSession(userID)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("session:%s", session.ID)
	err = s.client.Set(key, strconv.FormatInt(int64(userID), 10), SessionExpiration).Err()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return session, nil
}
