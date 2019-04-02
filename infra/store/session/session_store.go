package sessionstore

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type sessionStoreImpl struct {
	ctx    context.Context
	client *redis.Client
	db     *sqlutil.DB
}

const (
	// SessionExpiration is sessions expiration
	SessionExpiration = time.Hour * 24 * 7

	keyPrefix = "session"
)

var (
	errSessionNotFound = fmt.Errorf("session not found")
)

// NewSessionStore returns new session store
func NewSessionStore(ctx context.Context, cli *redis.Client, db *sqlutil.DB) store.SessionStore {
	return &sessionStoreImpl{
		ctx:    ctx,
		client: cli,
		db:     db,
	}
}

func (s *sessionStoreImpl) CreateSession(nameOrEmail, password string) (*model.Session, error) {
	u, err := record.Users(qm.Where("email = ? or name = ?", nameOrEmail, nameOrEmail)).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.WithStack(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	session, err := s.setSession(model.UserID(u.ID))
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionStoreImpl) GetSession(sessionID string) (*model.Session, error) {
	keys, err := s.client.Keys(redisKey(sessionID) + ":*").Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(keys) == 0 {
		return nil, errors.WithStack(errSessionNotFound)
	}

	v, err := s.client.Get(keys[0]).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	id, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.Session{
		ID:     sessionID,
		UserID: model.UserID(id),
	}, nil
}

func (s *sessionStoreImpl) DeleteSession(sessionID string) error {
	keys, err := s.client.Keys(redisKey(sessionID) + ":*").Result()
	if err != nil {
		return errors.WithStack(err)
	}

	if len(keys) == 0 {
		return errors.WithStack(errSessionNotFound)
	}

	_, err = s.client.Del(keys...).Result()
	return errors.WithStack(err)
}

func (s *sessionStoreImpl) ResetSession(userID model.UserID) (*model.Session, error) {
	keys, err := s.client.Keys(fmt.Sprintf("%s:*:%v", keyPrefix, userID)).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	_, err = s.client.Del(keys...).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	session, err := s.setSession(userID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionStoreImpl) setSession(userID model.UserID) (*model.Session, error) {
	session, err := model.NewSession(userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	id := strconv.FormatInt(int64(userID), 10)
	key := redisKey(session.ID) + ":" + id
	err = s.client.Set(key, id, SessionExpiration).Err()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return session, nil
}

func redisKey(sessionID string) string {
	return fmt.Sprintf("%s:%s", keyPrefix, sessionID)
}
