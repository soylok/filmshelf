package movies

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/params"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

func (s *MovieService) Create(ctx context.Context, movieParams params.MovieCreateParams) (domain.Movie, error) {
	movieParams.Title = strings.Join(strings.Fields(movieParams.Title), " ")
	if movieParams.Description != nil {
		normalized := strings.Join(strings.Fields(*movieParams.Description), " ")
		if len(normalized) == 0 {
			movieParams.Description = nil
		} else {
			movieParams.Description = &normalized
		}
	}

	movie := domain.Movie{
		ID:                 domain.UninitializedID,
		Title:              movieParams.Title,
		Description:        movieParams.Description,
		ReleaseYear:        movieParams.ReleaseYear,
		Rating:             movieParams.Rating,
		Budget:             movieParams.Budget,
		WorldWideBoxOffice: movieParams.WorldWideBoxOffice,
		DurationMinutes:    movieParams.DurationMinutes,
		Version:            domain.UninitializedVersion,
	}
	if err := movie.Validate(); err != nil {
		return domain.Movie{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}
	if err := movieParams.Validate(); err != nil {
		return domain.Movie{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}
	createdMovie, err := s.repo.Create(ctx, movieParams)
	if err != nil {
		var refErr *outbound.ReferenceNotFoundError
		if errors.As(err, &refErr) {
			return domain.Movie{}, &apperrors.Error{
				Type:    apperrors.ErrInvalidArgument,
				Message: fmt.Sprintf("one or more specified %s IDs do not exist", refErr.Entity),
				Cause:   err,
			}
		}
		return domain.Movie{}, err
	}
	return createdMovie, nil
}
