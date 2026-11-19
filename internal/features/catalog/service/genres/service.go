package genres

import "github.com/soylok/filmshelf/internal/features/catalog/ports/out"

type GenreService struct {
	repo outbound.GenreRepository
}

func NewGenreService(repository outbound.GenreRepository) *GenreService {
	return &GenreService{
		repo: repository,
	}
}
