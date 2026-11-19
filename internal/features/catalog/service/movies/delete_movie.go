package movies

import (
	"context"
	"errors"
	"fmt"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

func (s *MovieService) Delete(ctx context.Context, id int64) error {
	filename, err := s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, outbound.ErrNotFound) {
			return &apperrors.Error{
				Type:    apperrors.ErrNotFound,
				Message: "movie not found",
				Cause:   err,
			}
		}
		return err
	}
	if filename != nil {
		if err := s.posterStorage.Delete(*filename); err != nil {
			return fmt.Errorf("delete movie poster: %w", err)
		}
	}
	return nil
}
