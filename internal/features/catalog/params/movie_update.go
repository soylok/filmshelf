package params

import (
	"errors"

	"github.com/soylok/filmshelf/internal/core/domain"
)

type MovieUpdateParams struct {
	Title              domain.Nullable[string]
	Description        domain.Nullable[string]
	ReleaseYear        domain.Nullable[int32]
	Rating             domain.Nullable[float64]
	Budget             domain.Nullable[int64]
	WorldWideBoxOffice domain.Nullable[int64]
	DurationMinutes    domain.Nullable[int32]
	Countries          domain.Nullable[[]int64]
	Directors          domain.Nullable[[]int64]
	Genres             domain.Nullable[[]int64]
}

func (m MovieUpdateParams) Validate() error {
	if m.Title.Set && m.Title.Value == nil {
		return errors.New("title cannot be empty")
	}
	if m.Countries.Set {
		if m.Countries.Value == nil {
			return errors.New("countries must not be null")
		}
		if !isUnique(*m.Countries.Value) {
			return errors.New("country IDs must be unique")
		}
	}
	if m.Directors.Set {
		if m.Directors.Value == nil {
			return errors.New("directors must not be null")
		}
		if !isUnique(*m.Directors.Value) {
			return errors.New("director IDs must be unique")
		}
	}
	if m.Genres.Set {
		if m.Genres.Value == nil {
			return errors.New("genres must not be null")
		}
		if !isUnique(*m.Genres.Value) {
			return errors.New("genre IDs must be unique")
		}
	}
	return nil
}

func (m MovieUpdateParams) ApplyTo(movie *domain.Movie) error {
	tmp := *movie
	if m.Title.Set && m.Title.Value != nil {
		tmp.Title = *m.Title.Value
	}
	if m.Description.Set {
		tmp.Description = m.Description.Value
	}
	if m.ReleaseYear.Set {
		tmp.ReleaseYear = m.ReleaseYear.Value
	}
	if m.Rating.Set {
		tmp.Rating = m.Rating.Value
	}
	if m.Budget.Set {
		tmp.Budget = m.Budget.Value
	}
	if m.WorldWideBoxOffice.Set {
		tmp.WorldWideBoxOffice = m.WorldWideBoxOffice.Value
	}
	if m.DurationMinutes.Set {
		tmp.DurationMinutes = m.DurationMinutes.Value
	}

	if err := tmp.Validate(); err != nil {
		return err
	}
	*movie = tmp
	return nil
}
