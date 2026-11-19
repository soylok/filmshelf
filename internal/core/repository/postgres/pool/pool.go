package pool

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
}

func NewPool(ctx context.Context, config Config) (*Pool, error) {
	pgxcfg, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("parse pgx config: %w", err)
	}
	pgxcfg.ConnConfig.User = config.User
	pgxcfg.ConnConfig.Password = config.Password
	pgxcfg.ConnConfig.Host = config.Host
	pgxcfg.ConnConfig.Port = config.Port
	pgxcfg.ConnConfig.Database = config.DB
	pgxcfg.ConnConfig.Config.TLSConfig = nil
	pgxcfg.ConnConfig.Config.Fallbacks = nil

	pool, err := pgxpool.NewWithConfig(ctx, pgxcfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return &Pool{
		Pool: pool,
	}, nil
}
