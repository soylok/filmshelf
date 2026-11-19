package outbound

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
)

type CountryRepository interface {
	List(ctx context.Context) ([]domain.Country, error)
}
