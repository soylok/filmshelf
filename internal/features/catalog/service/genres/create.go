package genres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/domain"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

func (g *GenreService) Create(ctx context.Context, genre domain.Genre) (domain.Genre, error) {
	genre.Name = strings.Join(strings.Fields(genre.Name), " ")

	if err := genre.Validate(); err != nil {
		return domain.Genre{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}
	createdGenre, err := g.repo.Create(ctx, genre)
	if err != nil {
		if errors.Is(err, outbound.ErrAlreadyExists) {
			return domain.Genre{}, &apperrors.Error{
				Type:    apperrors.ErrConflict,
				Message: fmt.Sprintf("genre with name %q already exists", genre.Name),
				Cause:   err,
			}
		}
		return domain.Genre{}, err
	}
	return createdGenre, nil
}
