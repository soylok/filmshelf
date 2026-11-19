package countries

import (
	"time"

	"github.com/soylok/filmshelf/internal/core/repository/postgres"
)

type CountryRepository struct {
	querier   postgres.Querier
	opTimeout time.Duration
}

func NewCountryRepository(querier postgres.Querier, opTimeout time.Duration) *CountryRepository {
	return &CountryRepository{
		querier:   querier,
		opTimeout: opTimeout,
	}
}
