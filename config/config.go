package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	POSTGRES_HOST           string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT           string `mapstructure:"POSTGRES_PORT"`
	POSTGRES_DB             string `mapstructure:"POSTGRES_DB"`
	POSTGRES_USER           string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD       string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_SSL_MODE       string `mapstructure:"POSTGRES_SSL_MODE"`
	REDIS_HOST              string `mapstructure:"REDIS_HOST"`
	REDIS_PORT              string `mapstructure:"REDIS_PORT"`
	REDIS_PASSWORD          string `mapstructure:"REDIS_PASSWORD"`
	MINIO_ENDPOINT          string `mapstructure:"MINIO_ENDPOINT"`
	MINIO_PORT              string `mapstructure:"MINIO_PORT"`
	MINIO_ACCESS_KEY_ID     string `mapstructure:"MINIO_ACCESS_KEY_ID"`
	MINIO_SECRET_ACCESS_KEY string `mapstructure:"MINIO_SECRET_ACCESS_KEY"`
	MINIO_USE_SSL           string `mapstructure:"MINIO_USE_SSL"`
	OIDC_USERINFO_ENDPOINT  string `mapstructure:"OIDC_USERINFO_ENDPOINT"`
	STORAGE_BASE_URL        string `mapstructure:"STORAGE_BASE_URL"`
	STORAGE_DEFAULT_BUCKET  string `mapstructure:"STORAGE_DEFAULT_BUCKET"`
}

// LoadConfig loads the application configuration from environment variables or a configuration file.
// It returns the loaded configuration and an error if any required configuration values are missing or if there was an error loading the configuration.
func LoadConfig() (config Config, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		// Load configuration from environment variables in production environment.
		return Config{
			POSTGRES_HOST:           os.Getenv("POSTGRES_HOST"),
			POSTGRES_PORT:           os.Getenv("POSTGRES_PORT"),
			POSTGRES_DB:             os.Getenv("POSTGRES_DB"),
			POSTGRES_USER:           os.Getenv("POSTGRES_USER"),
			POSTGRES_PASSWORD:       os.Getenv("POSTGRES_PASSWORD"),
			POSTGRES_SSL_MODE:       os.Getenv("POSTGRES_SSL_MODE"),
			REDIS_HOST:              os.Getenv("REDIS_HOST"),
			REDIS_PORT:              os.Getenv("REDIS_PORT"),
			REDIS_PASSWORD:          os.Getenv("REDIS_PASSWORD"),
			MINIO_ENDPOINT:          os.Getenv("MINIO_ENDPOINT"),
			MINIO_PORT:              os.Getenv("MINIO_PORT"),
			MINIO_ACCESS_KEY_ID:     os.Getenv("MINIO_ACCESS_KEY_ID"),
			MINIO_SECRET_ACCESS_KEY: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
			MINIO_USE_SSL:           os.Getenv("MINIO_USE_SSL"),
			OIDC_USERINFO_ENDPOINT:  os.Getenv("OIDC_USERINFO_ENDPOINT"),
			STORAGE_BASE_URL:        os.Getenv("STORAGE_BASE_URL"),
			STORAGE_DEFAULT_BUCKET:  os.Getenv("STORAGE_DEFAULT_BUCKET"),
		}, nil
	}

	// Load configuration from a configuration file in other environments.
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// Check if required configuration values are missing.
	if config.POSTGRES_HOST == "" {
		err = errors.New("POSTGRES_HOST is required")
	}

	if config.POSTGRES_PORT == "" {
		err = errors.New("POSTGRES_PORT is required")
	}

	if config.POSTGRES_DB == "" {
		err = errors.New("POSTGRES_DB is required")
	}

	if config.POSTGRES_USER == "" {
		err = errors.New("POSTGRES_USER is required")
	}

	if config.POSTGRES_PASSWORD == "" {
		err = errors.New("POSTGRES_PASSWORD is required")
	}

	if config.POSTGRES_SSL_MODE == "" {
		err = errors.New("POSTGRES_SSL_MODE is required")
	}

	if config.REDIS_HOST == "" {
		err = errors.New("REDIS_HOST is required")
	}

	if config.REDIS_PORT == "" {
		err = errors.New("REDIS_PORT is required")
	}

	if config.REDIS_PASSWORD == "" {
		err = errors.New("REDIS_PASSWORD is required")
	}

	if config.MINIO_ENDPOINT == "" {
		err = errors.New("MINIO_ENDPOINT is required")
	}

	if config.MINIO_PORT == "" {
		err = errors.New("MINIO_PORT is required")
	}

	if config.MINIO_ACCESS_KEY_ID == "" {
		err = errors.New("MINIO_ACCESS_KEY_ID is required")
	}
	if config.MINIO_SECRET_ACCESS_KEY == "" {
		err = errors.New("MINIO_SECRET_ACCESS_KEY is required")
	}
	if config.MINIO_USE_SSL == "" {
		err = errors.New("MINIO_USE_SSL is required")
	}

	if config.OIDC_USERINFO_ENDPOINT == "" {
		err = errors.New("OIDC_USERINFO_ENDPOINT is required")
	}

	if config.STORAGE_BASE_URL == "" {
		err = errors.New("STORAGE_BASE_URL is required")
	}

	if config.STORAGE_DEFAULT_BUCKET == "" {
		err = errors.New("STORAGE_DEFAULT_BUCKET is required")
	}

	return
}
