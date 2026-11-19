package dto

import "github.com/soylok/filmshelf/internal/core/domain"

type CountryDTO struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
}

func ToCountryDTO(country domain.Country) CountryDTO {
	return CountryDTO{
		ID:   country.ID,
		Code: country.Code,
	}
}

func CountriesToDTO(countries []domain.Country) []CountryDTO {
	dtos := make([]CountryDTO, 0, len(countries))
	for _, country := range countries {
		dtos = append(dtos, ToCountryDTO(country))
	}
	return dtos
}
