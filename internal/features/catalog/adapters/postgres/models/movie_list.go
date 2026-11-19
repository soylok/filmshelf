package models

import "github.com/soylok/filmshelf/internal/features/catalog/ports"

type MovieListItemModel struct {
	ID             int64    `db:"id"`
	Title          string   `db:"title"`
	ReleaseYear    *int32   `db:"release_year"`
	Rating         *float64 `db:"rating"`
	PosterFilename *string  `db:"poster_filename"`
	Genre          *string  `db:"genre"`
}

func (m MovieListItemModel) ToListItem() ports.MovieListItem {
	return ports.MovieListItem{
		ID:             m.ID,
		Title:          m.Title,
		ReleaseYear:    m.ReleaseYear,
		Rating:         m.Rating,
		Genre:          m.Genre,
		PosterFilename: m.PosterFilename,
	}
}

func MoviesToDomain(models []MovieListItemModel) []ports.MovieListItem {
	movies := make([]ports.MovieListItem, 0, len(models))
	for _, model := range models {
		movies = append(movies, model.ToListItem())
	}
	return movies
}
