package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config represents app config
type Config struct {
	DataBaseURL     string `envconfig:"database_url" required:"true"`
	TestDataBaseURL string `envconfig:"test_database_url" required:"true"`
	RedisAddr       string `envconfig:"redis_addr" required:"true"`
	DebugLog        bool   `envconfig:"debug_log"`
	JobIntervalSec  int    `envconfig:"job_interval_sec" required:"true"`
	HydraAdminURL   string `envconfig:"hydra_admin_url" required:"true"`
	MinioPublicURL  string `envconfig:"minio_public_url" required:"true"`
	MinioEndpoint   string `envconfig:"minio_endpoint" required:"true"`
	MinioAccessKey  string `envconfig:"minio_access_key" required:"true"`
	MinioSecretKey  string `envconfig:"minio_secret_key" required:"true"`
	MinioBucketName string `envconfig:"minio_bucket_name" required:"true"`
}

// LoadConfig loads config
func LoadConfig(filenames ...string) (*Config, error) {
	// do not care if .env does not exist.
	_ = godotenv.Overload(filenames...)

	c := &Config{}
	err := envconfig.Process("", c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}
	return c, nil
}
