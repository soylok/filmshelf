package directors

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
)

func (s *DirectorService) List(ctx context.Context) ([]domain.Director, error) {
	directors, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return directors, nil
}
