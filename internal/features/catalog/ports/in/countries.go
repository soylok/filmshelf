package inbound

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
)

type CountryService interface {
	List(ctx context.Context) ([]domain.Country, error)
}
