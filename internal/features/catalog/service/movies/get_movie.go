package movies

import (
	"context"
	"errors"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/domain"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

func (s *MovieService) Get(ctx context.Context, id int64) (domain.Movie, error) {
	movie, err := s.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, outbound.ErrNotFound) {
			return domain.Movie{}, &apperrors.Error{
				Type:    apperrors.ErrNotFound,
				Message: "movie not found",
				Cause:   err,
			}
		}
		return domain.Movie{}, err
	}
	return movie, nil
}
