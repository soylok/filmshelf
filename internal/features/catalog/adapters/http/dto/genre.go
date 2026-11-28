package dto

import "github.com/soylok/filmshelf/internal/core/domain"

type GenreDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GenreCreateRequest struct {
	Name string `json:"name"`
}

func GenreToDTO(genre domain.Genre) GenreDTO {
	return GenreDTO{
		ID:   genre.ID,
		Name: genre.Name,
	}
}

func GenresToDTO(genres []domain.Genre) []GenreDTO {
	dtos := make([]GenreDTO, 0, len(genres))
	for _, genre := range genres {
		dtos = append(dtos, GenreToDTO(genre))
	}
	return dtos
}
