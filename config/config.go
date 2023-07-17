package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	HOST                             string `mapstructure:"HOST"`
	PORT                             string `mapstructure:"PORT"`
	POSTGRES_DB                      string `mapstructure:"POSTGRES_DB"`
	POSTGRES_USER                    string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD                string `mapstructure:"POSTGRES_PASSWORD"`
	SSL_MODE                         string `mapstructure:"SSL_MODE"`
	OPENID_CONNECT_USERINFO_ENDPOINT string `mapstructure:"OPENID_CONNECT_USERINFO_ENDPOINT"`
}

// LoadConfig loads the application configuration from environment variables or a configuration file.
// It returns the loaded configuration and an error if any required configuration values are missing or if there was an error loading the configuration.
func LoadConfig() (config Config, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		// Load configuration from environment variables in production environment.
		return Config{
			HOST:                             os.Getenv("HOST"),
			PORT:                             os.Getenv("PORT"),
			POSTGRES_DB:                      os.Getenv("POSTGRES_DB"),
			POSTGRES_USER:                    os.Getenv("POSTGRES_USER"),
			POSTGRES_PASSWORD:                os.Getenv("POSTGRES_PASSWORD"),
			SSL_MODE:                         os.Getenv("SSL_MODE"),
			OPENID_CONNECT_USERINFO_ENDPOINT: os.Getenv("OPENID_CONNECT_USERINFO_ENDPOINT"),
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
	if config.HOST == "" {
		err = errors.New("HOST is required")
	}

	if config.PORT == "" {
		err = errors.New("PORT is required")
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

	if config.SSL_MODE == "" {
		err = errors.New("SSL_MODE is required")
	}

	if config.OPENID_CONNECT_USERINFO_ENDPOINT == "" {
		err = errors.New("OPENID_CONNECT_USERINFO_ENDPOINT is required")
	}

	return
}
