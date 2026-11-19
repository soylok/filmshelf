package genres

import (
	"context"
	"errors"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

func (s *GenreService) Delete(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, outbound.ErrNotFound) {
			return &apperrors.Error{
				Type:    apperrors.ErrNotFound,
				Message: "genre not found",
				Cause:   err,
			}
		}
		return err
	}
	return nil
}
