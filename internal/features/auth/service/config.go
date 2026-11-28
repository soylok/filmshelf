package auth

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SessionTTL             time.Duration `envconfig:"SESSION_TTL" default:"720h"`
	SessionCleanupInterval time.Duration `envconfig:"SESSION_CLEANUP_INTERVAL" default:"1h"`
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("AUTH", &cfg); err != nil {
		return Config{}, fmt.Errorf("parse auth config: %w", err)
	}
	if cfg.SessionTTL <= 0 {
		return Config{}, fmt.Errorf("AUTH_SESSION_TTL must be greater than zero")
	}
	if cfg.SessionCleanupInterval < time.Minute {
		return Config{}, fmt.Errorf("AUTH_SESSION_CLEANUP_INTERVAL must be at least 1 minute")
	}
	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
