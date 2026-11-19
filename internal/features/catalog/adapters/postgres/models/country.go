package models

import "github.com/soylok/filmshelf/internal/core/domain"

type CountryModel struct {
	ID   int64  `db:"id"`
	Code string `db:"code"`
}

func (c CountryModel) ToDomain() domain.Country {
	return domain.Country{
		ID:   c.ID,
		Code: c.Code,
	}
}

func CountriesToDomain(models []CountryModel) []domain.Country {
	countries := make([]domain.Country, 0, len(models))
	for _, model := range models {
		countries = append(countries, model.ToDomain())
	}
	return countries
}
