package pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	User      string        `envconfig:"USER" required:"true"`
	Password  string        `envconfig:"PASSWORD" required:"true"`
	Host      string        `envconfig:"HOST" default:"localhost"`
	Port      uint16           `envconfig:"PORT" default:"5432"`
	DB        string        `envconfig:"DB" required:"true"`
	OpTimeout time.Duration `envconfig:"OP_TIMEOUT" default:"5s"`
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("POSTGRES", &cfg); err != nil {
		return Config{}, fmt.Errorf("parsing pool config: %w", err)
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
