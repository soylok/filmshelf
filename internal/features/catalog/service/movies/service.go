package movies

import outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"

type MovieService struct {
	repo          outbound.MovieRepository
	posterStorage outbound.PosterStorage
}

func NewMovieService(repository outbound.MovieRepository, posterStorage outbound.PosterStorage) *MovieService {
	return &MovieService{
		repo:          repository,
		posterStorage: posterStorage,
	}
}
