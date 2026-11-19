package params

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
	"unicode/utf8"
)

type MovieListParams struct {
	Limit       *int64
	Offset      *int64
	Q           *string
	GenreIDs    []int64
	CountryIDs  []int64
	DirectorIDs []int64
	Year        *int32
	Rating      *float64
	Sort        *MovieSort
}

func (p MovieListParams) Validate() error {
	if p.Limit != nil {
		if *p.Limit <= 0 {
			return errors.New("limit must be greater than 0")
		}
		if *p.Limit > 100 {
			return errors.New("limit must not exceed 100")
		}
	}

	if p.Offset != nil && *p.Offset < 0 {
		return errors.New("offset must not fall below 0")
	}

	if p.Q != nil {
		if strings.TrimSpace(*p.Q) == "" {
			return errors.New("search query must not be empty")
		}
		if utf8.RuneCountInString(*p.Q) > 255 {
			return errors.New("search query must not exceed 255 chars")
		}
	}

	if !isUnique(p.GenreIDs) {
		return errors.New("genre IDs must be unique")
	}

	if !isUnique(p.CountryIDs) {
		return errors.New("country IDs must be unique")
	}

	if !isUnique(p.DirectorIDs) {
		return errors.New("director IDs must be unique")
	}

	if p.Year != nil {
		if *p.Year < 1800 {
			return errors.New("release year must be at least 1800")
		}
		if maxReleaseYear := int32(time.Now().Year()) + 5; *p.Year > maxReleaseYear {
			return fmt.Errorf("release year must not exceed %d", maxReleaseYear)
		}
	}

	if p.Rating != nil {
		if math.IsNaN(*p.Rating) || math.IsInf(*p.Rating, 0) {
			return errors.New("rating must be a finite number")
		}
		if *p.Rating < 0 {
			return errors.New("rating must not fall below 0")
		}
		if *p.Rating > 10 {
			return errors.New("rating must not exceed 10")
		}
	}

	if p.Sort != nil {
		switch *p.Sort {
		case MovieSortIDAsc, MovieSortIDDesc, MovieSortTitleAsc, MovieSortTitleDesc,
			MovieSortYearAsc, MovieSortYearDesc, MovieSortRatingAsc, MovieSortRatingDesc:
		default:
			return fmt.Errorf("unsupported movie sort %q", *p.Sort)
		}
	}

	return nil
}
