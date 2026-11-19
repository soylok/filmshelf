package countries

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
)

func (s *CountryService) List(ctx context.Context) ([]domain.Country, error) {
	countries, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
