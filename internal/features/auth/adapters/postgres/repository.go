package postgres

import (
	"time"

	"github.com/soylok/filmshelf/internal/core/repository/postgres"
)

type AuthRepository struct {
	querier   postgres.Querier
	opTimeout time.Duration
}

func NewAuthRepository(querier postgres.Querier, opTimeout time.Duration) *AuthRepository {
	return &AuthRepository{
		querier:   querier,
		opTimeout: opTimeout,
	}
}
