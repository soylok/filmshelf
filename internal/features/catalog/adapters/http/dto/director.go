package dto

import "github.com/soylok/filmshelf/internal/core/domain"

type DirectorDTO struct {
	ID         int64   `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
}

type DirectorCreateRequest struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
}

func DirectorToDTO(director domain.Director) DirectorDTO {
	return DirectorDTO{
		ID:         director.ID,
		FirstName:  director.FirstName,
		LastName:   director.LastName,
		MiddleName: director.MiddleName,
	}
}

func DirectorsToDTO(directors []domain.Director) []DirectorDTO {
	dtos := make([]DirectorDTO, 0, len(directors))
	for _, director := range directors {
		dtos = append(dtos, DirectorToDTO(director))
	}
	return dtos
}
