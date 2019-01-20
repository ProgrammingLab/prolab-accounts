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

	"github.com/ProgrammingLab/prolab-accounts/dao"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/store"
)

type sessionStoreImpl struct {
	ctx    context.Context
	client *redis.Client
	db     *sql.DB
}

const (
	// SessionExpiration is sessions expiration
	SessionExpiration = time.Hour * 24 * 7
)

// NewSessionStore returns new session store
func NewSessionStore(ctx context.Context, cli *redis.Client, db *sql.DB) store.SessionStore {
	return &sessionStoreImpl{
		ctx:    ctx,
		client: cli,
		db:     db,
	}
}

func (s *sessionStoreImpl) CreateSession(nameOrEmail, password string) (*model.Session, error) {
	u, err := dao.Users(qm.Where("email = ? or name = ?", nameOrEmail, nameOrEmail)).One(s.ctx, s.db)
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

	session, err := model.NewSession(model.UserID(u.ID))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	key := redisKey(session.ID)
	err = s.client.Set(key, strconv.FormatInt(u.ID, 10), SessionExpiration).Err()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return session, nil
}

func (s *sessionStoreImpl) GetSession(sessionID string) (*model.Session, error) {
	v, err := s.client.Get(redisKey(sessionID)).Result()
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

func redisKey(sessionID string) string {
	return fmt.Sprintf("session:%s", sessionID)
}
