package domain

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

type Movie struct {
	ID                 int64
	Title              string
	Description        *string
	ReleaseYear        *int32
	Rating             *float64
	Budget             *int64
	WorldWideBoxOffice *int64
	DurationMinutes    *int32
	PosterFilename     *string
	Countries          []Country
	Directors          []Director
	Genres             []Genre
	Version            int64
}

func (m Movie) Validate() error {
	switch titleLen := utf8.RuneCountInString(m.Title); {
	case titleLen < 1:
		return errors.New("title must not be empty")
	case titleLen > 255:
		return errors.New("title must not exceed 255 chars")
	}

	if m.ReleaseYear != nil {
		if *m.ReleaseYear < 1800 {
			return errors.New("release year must be at least 1800")
		}
		if maxReleaseYear := int32(time.Now().Year()) + 5; *m.ReleaseYear > maxReleaseYear {
			return fmt.Errorf("release year must not exceed %d", maxReleaseYear)
		}
	}

	if m.Rating != nil {
		if *m.Rating < 0 {
			return errors.New("rating must not fall below 0")
		}
		if *m.Rating > 10 {
			return errors.New("rating must not exceed 10")
		}
	}

	if m.Budget != nil && *m.Budget < 0 {
		return errors.New("budget must not fall below 0")
	}

	if m.WorldWideBoxOffice != nil && *m.WorldWideBoxOffice < 0 {
		return errors.New("world wide box office must not fall below 0")
	}

	if m.DurationMinutes != nil && *m.DurationMinutes <= 0 {
		return errors.New("duration must be greater than 0")
	}
	return nil
}
