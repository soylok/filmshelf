package genres

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
)

func (s *GenreService) List(ctx context.Context) ([]domain.Genre, error) {
	genres, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return genres, nil
}
