package countries

import (
	"context"
	"fmt"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/models"
)

func (repo *CountryRepository) List(ctx context.Context) ([]domain.Country, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `select
		id,
		code
	from filmshelf.countries
	order by code`
	rows, err := repo.querier.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	countriesModel := make([]models.CountryModel, 0)
	for rows.Next() {
		var countryModel models.CountryModel
		err := rows.Scan(
			&countryModel.ID,
			&countryModel.Code,
		)
		if err != nil {
			return nil, fmt.Errorf("scan country: %w", err)
		}
		countriesModel = append(countriesModel, countryModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate countries: %w", err)
	}

	return models.CountriesToDomain(countriesModel), nil
}
