package models

import "github.com/soylok/filmshelf/internal/core/domain"

type GenreModel struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Version int64  `db:"version"`
}

func (g GenreModel) ToDomain() domain.Genre {
	return domain.Genre{
		ID:      g.ID,
		Name:    g.Name,
		Version: g.Version,
	}
}

func GenresToDomain(models []GenreModel) []domain.Genre {
	genres := make([]domain.Genre, 0, len(models))
	for _, model := range models {
		genres = append(genres, model.ToDomain())
	}
	return genres
}
