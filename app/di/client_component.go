package di

import (
	"context"

	minio "github.com/minio/minio-go"
	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
)

// ClientComponent is an interface of api clients
type ClientComponent interface {
	HydraClient(ctx context.Context) *hydra.CodeGenSDK
	MinioClient(ctx context.Context) *minio.Client
}

// NewClientComponent returns new client component
func NewClientComponent(cfg *config.Config) (ClientComponent, error) {
	h, err := newHydraClient(cfg)
	if err != nil {
		return nil, err
	}

	m, err := newMinioClient(cfg)
	if err != nil {
		return nil, err
	}

	return &clientComponentImpl{
		hydraCli: h,
		minioCli: m,
	}, nil
}

func newHydraClient(cfg *config.Config) (*hydra.CodeGenSDK, error) {
	hc := &hydra.Configuration{
		AdminURL: cfg.HydraAdminURL,
	}
	cli, err := hydra.NewSDK(hc)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return cli, nil
}

func newMinioClient(cfg *config.Config) (*minio.Client, error) {
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

type clientComponentImpl struct {
	hydraCli *hydra.CodeGenSDK
	minioCli *minio.Client
}

func (c *clientComponentImpl) HydraClient(ctx context.Context) *hydra.CodeGenSDK {
	return c.hydraCli
}

func (c *clientComponentImpl) MinioClient(ctx context.Context) *minio.Client {
	return c.minioCli
}
