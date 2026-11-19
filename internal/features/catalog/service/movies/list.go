package movies

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/features/catalog/params"
	"github.com/soylok/filmshelf/internal/features/catalog/ports"
)

const (
	defaultMovieListLimit  int64 = 20
	defaultMovieListOffset int64 = 0
	defaultMovieListSort         = params.MovieSortIDDesc
)

func (s *MovieService) List(ctx context.Context, movieParams params.MovieListParams) ([]ports.MovieListItem, error) {
	if movieParams.Limit == nil {
		limit := defaultMovieListLimit
		movieParams.Limit = &limit
	}
	if movieParams.Offset == nil {
		offset := defaultMovieListOffset
		movieParams.Offset = &offset
	}
	if movieParams.Sort == nil {
		sort := defaultMovieListSort
		movieParams.Sort = &sort
	}

	if err := movieParams.Validate(); err != nil {
		return nil, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}
	movies, err := s.repo.List(ctx, movieParams)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
