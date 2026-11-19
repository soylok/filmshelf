package directors

import (
	"time"

	"github.com/soylok/filmshelf/internal/core/repository/postgres"
)

type DirectorRepository struct {
	querier   postgres.Querier
	opTimeout time.Duration
}

func NewDirectorRepository(querier postgres.Querier, opTimeout time.Duration) *DirectorRepository {
	return &DirectorRepository{
		querier:   querier,
		opTimeout: opTimeout,
	}
}
