package models

import "github.com/soylok/filmshelf/internal/core/domain"

type MovieModel struct {
	ID                 int64           `db:"id"`
	Title              string          `db:"title"`
	Description        *string         `db:"description"`
	ReleaseYear        *int32          `db:"release_year"`
	Rating             *float64        `db:"rating"`
	Budget             *int64          `db:"budget"`
	WorldWideBoxOffice *int64          `db:"worldwide_box_office"`
	DurationMinutes    *int32          `db:"duration_minutes"`
	PosterFilename     *string         `db:"poster_filename"`
	Countries          []CountryModel  `db:"countries"`
	Directors          []DirectorModel `db:"directors"`
	Genres             []GenreModel    `db:"genres"`
	Version            int64           `db:"version"`
}

func (m MovieModel) ToDomain() domain.Movie {
	return domain.Movie{
		ID:                 m.ID,
		Title:              m.Title,
		Description:        m.Description,
		ReleaseYear:        m.ReleaseYear,
		Rating:             m.Rating,
		Budget:             m.Budget,
		WorldWideBoxOffice: m.WorldWideBoxOffice,
		DurationMinutes:    m.DurationMinutes,
		PosterFilename:     m.PosterFilename,
		Countries:          CountriesToDomain(m.Countries),
		Directors:          DirectorsToDomain(m.Directors),
		Genres:             GenresToDomain(m.Genres),
		Version:            m.Version,
	}
}
