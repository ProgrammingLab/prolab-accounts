package userstore

import (
	"bytes"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"image"
	_ "image/gif" // for image
	_ "image/jpeg"
	_ "image/png"

	minio "github.com/minio/minio-go"
	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type userStoreImpl struct {
	ctx        context.Context
	db         *sqlutil.DB
	cli        *minio.Client
	bucketName string
}

// NewUserStore returns new user store
func NewUserStore(ctx context.Context, db *sqlutil.DB, cli *minio.Client, bucket string) store.UserStore {
	return &userStoreImpl{
		ctx:        ctx,
		db:         db,
		cli:        cli,
		bucketName: bucket,
	}
}

func (s *userStoreImpl) CreateUser(user *record.User) error {
	err := user.Insert(s.ctx, s.db, boil.Infer())
	return errors.WithStack(err)
}

func (s *userStoreImpl) GetPublicUserByName(name string) (*record.User, error) {
	mods := []qm.QueryMod{
		qm.Load("Profile", record.ProfileWhere.ProfileScope.EQ(null.IntFrom(int(model.Public)))),
		qm.Load("Profile.Role"),
		qm.Load("Profile.Department"),
		record.UserWhere.Name.EQ(name),
	}
	u, err := record.Users(mods...).One(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return u, nil
}

func (s *userStoreImpl) GetUserByName(name string) (*record.User, error) {
	mods := []qm.QueryMod{
		qm.Load("Profile.Role"),
		qm.Load("Profile.Department"),
		record.UserWhere.Name.EQ(name),
	}
	u, err := record.Users(mods...).One(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return u, nil
}

func (s *userStoreImpl) GetUserByEmail(email string) (*record.User, error) {
	mods := []qm.QueryMod{
		qm.Load("Profile.Role"),
		qm.Load("Profile.Department"),
		record.UserWhere.Email.EQ(email),
	}
	u, err := record.Users(mods...).One(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return u, nil
}

func (s *userStoreImpl) GetUserWithPrivate(userID model.UserID) (*record.User, error) {
	mods := []qm.QueryMod{
		qm.Load("Profile.Role"),
		qm.Load("Profile.Department"),
		qm.Where("id = ?", userID),
	}
	u, err := record.Users(mods...).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "")
	}

	return u, nil
}

func (s *userStoreImpl) ListPublicUsers(minUserID model.UserID, limit int) ([]*record.User, model.UserID, error) {
	mods := []qm.QueryMod{
		qm.Load("Profile.Role"),
		qm.Load("Profile.Department"),
		qm.InnerJoin("profiles on profiles.id = users.profile_id"),
		qm.Where("? <= users.id", minUserID),
		qm.Where("profiles.profile_scope = ?", model.Public),
		qm.Limit(limit + 1),
		qm.OrderBy("users.id"),
	}

	u, err := record.Users(mods...).All(s.ctx, s.db)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	if len(u) <= limit {
		return u, 0, nil
	}
	return u[:limit], model.UserID(u[limit].ID), nil
}

func (s *userStoreImpl) ListPrivateUsers(minUserID model.UserID, limit int) ([]*record.User, model.UserID, error) {
	mods := []qm.QueryMod{
		qm.Load("Profile.Role"),
		qm.Load("Profile.Department"),
		qm.Where("? <= users.id", minUserID),
		qm.Limit(limit + 1),
		qm.OrderBy("users.id"),
	}

	u, err := record.Users(mods...).All(s.ctx, s.db)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	if len(u) <= limit {
		return u, 0, nil
	}
	return u[:limit], model.UserID(u[limit].ID), nil
}

func (s *userStoreImpl) UpdateFullName(userID model.UserID, fullName string) (*record.User, error) {
	var u *record.User
	err := s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		u, err = record.FindUser(s.ctx, tx, int64(userID))
		if err != nil {
			return errors.WithStack(err)
		}

		u.FullName = fullName
		_, err = u.Update(s.ctx, tx, boil.Whitelist(record.UserColumns.FullName, record.UserColumns.UpdatedAt))
		return errors.WithStack(err)
	})
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *userStoreImpl) UpdateIcon(userID model.UserID, icon []byte) (*record.User, error) {
	r := bytes.NewReader(icon)
	_, ext, err := image.DecodeConfig(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	name, err := generateFilename(ext)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	mods := []qm.QueryMod{
		qm.Load("Profile.Role"),
		qm.Load("Profile.Department"),
		qm.Where("id = ?", userID),
	}
	u, err := record.Users(mods...).One(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	opt := minio.PutObjectOptions{
		ContentType: "image/" + ext,
	}
	_, err = s.cli.PutObjectWithContext(s.ctx, s.bucketName, name, r, r.Size(), opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	old := u.AvatarFilename
	u.AvatarFilename = null.StringFrom(name)
	_, err = u.Update(s.ctx, s.db, boil.Whitelist(record.UserColumns.AvatarFilename, record.UserColumns.UpdatedAt))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !old.Valid {
		return u, nil
	}

	err = s.cli.RemoveObject(s.bucketName, old.String)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return u, nil
}

func generateFilename(ext string) (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.WithStack(err)
	}

	res := base64.RawURLEncoding.EncodeToString(b)

	return string(res) + "." + ext, nil
}

var selectQuery = map[model.ProfileScope]string{
	model.MembersOnly: "users.id, users.name, users.full_name, users.avatar_filename, users.profile_id",
	model.Public:      "users.id, users.name, users.avatar_filename, users.profile_id",
	model.Private:     "users.*",
}

func (s *userStoreImpl) selectQuery(scope model.ProfileScope) string {
	q, ok := selectQuery[scope]
	if !ok {
		return selectQuery[model.Public]
	}
	return q
}
