package movies

import (
	"time"

	"github.com/soylok/filmshelf/internal/core/repository/postgres"
)

type MovieRepository struct {
	querier   postgres.Querier
	opTimeout time.Duration
}

func NewMovieRepository(querier postgres.Querier, opTimeout time.Duration) *MovieRepository {
	return &MovieRepository{
		querier:   querier,
		opTimeout: opTimeout,
	}
}
