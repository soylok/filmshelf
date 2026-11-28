package models

import "github.com/soylok/filmshelf/internal/core/domain"

type GenreModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (g GenreModel) ToDomain() domain.Genre {
	return domain.Genre{
		ID:   g.ID,
		Name: g.Name,
	}
}

func GenresToDomain(models []GenreModel) []domain.Genre {
	genres := make([]domain.Genre, 0, len(models))
	for _, model := range models {
		genres = append(genres, model.ToDomain())
	}
	return genres
}
