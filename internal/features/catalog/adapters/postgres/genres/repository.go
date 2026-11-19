package genres

import (
	"time"

	"github.com/soylok/filmshelf/internal/core/repository/postgres"
)

type GenreRepository struct {
	querier   postgres.Querier
	opTimeout time.Duration
}

func NewGenreRepository(querier postgres.Querier, opTimeout time.Duration) *GenreRepository {
	return &GenreRepository{
		querier:   querier,
		opTimeout: opTimeout,
	}
}
