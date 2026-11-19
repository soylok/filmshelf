package media

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostersDir string `envconfig:"POSTERS_DIR" required:"true"`
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("MEDIA", &cfg); err != nil {
		return Config{}, fmt.Errorf("parse media config: %w", err)
	}
	if cfg.PostersDir == "" {
		return Config{}, fmt.Errorf("MEDIA_POSTERS_DIR must not be empty")
	}
	return cfg, nil
}
