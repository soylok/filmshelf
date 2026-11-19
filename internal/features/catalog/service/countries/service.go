package countries

import outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"

type CountryService struct {
	repo outbound.CountryRepository
}

func NewCountryService(repository outbound.CountryRepository) *CountryService {
	return &CountryService{
		repo: repository,
	}
}
