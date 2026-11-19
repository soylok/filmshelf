package models

import "github.com/soylok/filmshelf/internal/core/domain"

type DirectorModel struct {
	ID         int64   `db:"id"`
	FirstName  string  `db:"first_name"`
	LastName   string  `db:"last_name"`
	MiddleName *string `db:"middle_name"`
	Version    int64   `db:"version"`
}

func (d DirectorModel) ToDomain() domain.Director {
	return domain.Director{
		ID:         d.ID,
		FirstName:  d.FirstName,
		LastName:   d.LastName,
		MiddleName: d.MiddleName,
		Version:    d.Version,
	}
}

func DirectorsToDomain(models []DirectorModel) []domain.Director {
	directors := make([]domain.Director, 0, len(models))
	for _, model := range models {
		directors = append(directors, model.ToDomain())
	}
	return directors
}
