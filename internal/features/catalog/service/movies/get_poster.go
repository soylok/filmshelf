package movies

import (
	"context"
	"errors"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

func (s *MovieService) GetPosterFilename(ctx context.Context, movieID int64) (*string, error) {
	filename, err := s.repo.GetPosterFilename(ctx, movieID)
	if err != nil {
		if errors.Is(err, outbound.ErrNotFound) {
			return nil, &apperrors.Error{
				Type:    apperrors.ErrNotFound,
				Message: "film not found",
				Cause:   err,
			}
		}
		return nil, err
	}
	if filename == nil {
		return nil, &apperrors.Error{
			Type:    apperrors.ErrNotFound,
			Message: "poster not found",
			Cause:   apperrors.ErrNotFound,
		}
	}
	return filename, nil
}
