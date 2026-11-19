package directors

import outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"

type DirectorService struct {
	repo outbound.DirectorRepository
}

func NewDirectorService(repository outbound.DirectorRepository) *DirectorService {
	return &DirectorService{
		repo: repository,
	}
}
