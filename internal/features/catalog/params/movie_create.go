package params

import "errors"

type MovieCreateParams struct {
	Title              string
	Description        *string
	ReleaseYear        *int32
	Rating             *float64
	Budget             *int64
	WorldWideBoxOffice *int64
	DurationMinutes    *int32
	Countries          []int64
	Directors          []int64
	Genres             []int64
}

func (p MovieCreateParams) Validate() error {
	if !isUnique(p.Countries) {
		return errors.New("country IDs must be unique")
	}

	if !isUnique(p.Directors) {
		return errors.New("director IDs must be unique")
	}

	if !isUnique(p.Genres) {
		return errors.New("genre IDs must be unique")
	}
	return nil
}
