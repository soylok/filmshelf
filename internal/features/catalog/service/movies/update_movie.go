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

func (s *MovieService) Update(
	ctx context.Context,
	id int64,
	movieParams params.MovieUpdateParams,
) (domain.Movie, error) {
	if movieParams.Title.Set && movieParams.Title.Value != nil {
		normalized := strings.Join(strings.Fields(*movieParams.Title.Value), " ")
		if len(normalized) == 0 {
			movieParams.Title.Value = nil
		} else {
			movieParams.Title.Value = &normalized
		}
	}
	if movieParams.Description.Set && movieParams.Description.Value != nil {
		normalized := strings.Join(strings.Fields(*movieParams.Description.Value), " ")
		if len(normalized) == 0 {
			movieParams.Description.Value = nil
		} else {
			movieParams.Description.Value = &normalized
		}
	}

	if err := movieParams.Validate(); err != nil {
		return domain.Movie{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}

	movie, err := s.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, outbound.ErrNotFound) {
			return domain.Movie{}, &apperrors.Error{
				Type:    apperrors.ErrNotFound,
				Message: "movie not found",
				Cause:   err,
			}
		}
		return domain.Movie{}, fmt.Errorf("get movie for patch: %w", err)
	}
	if err := movieParams.ApplyTo(&movie); err != nil {
		return domain.Movie{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}
	movie, err = s.repo.Update(ctx, movie.ID, movie.Version, movieParams)
	if err != nil {
		var refErr *outbound.ReferenceNotFoundError
		if errors.As(err, &refErr) {
			return domain.Movie{}, &apperrors.Error{
				Type:    apperrors.ErrInvalidArgument,
				Message: fmt.Sprintf("one or more specified %s IDs do not exist", refErr.Entity),
				Cause:   err,
			}
		}
		if errors.Is(err, outbound.ErrConcurrentModification) {
			return domain.Movie{}, &apperrors.Error{
				Type:    apperrors.ErrConflict,
				Message: "movie was modified or deleted; reload and try again",
				Cause:   err,
			}
		}
		return domain.Movie{}, err
	}
	return movie, nil
}
