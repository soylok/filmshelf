package outbound

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/params"
	"github.com/soylok/filmshelf/internal/features/catalog/ports"
)

type MovieRepository interface {
	Get(ctx context.Context, id int64) (domain.Movie, error)
	GetPosterFilename(ctx context.Context, movieID int64) (*string, error)
	Delete(ctx context.Context, id int64) (*string, error)
	DeletePoster(ctx context.Context, id int64) (*string, error)
	Create(ctx context.Context, movieParams params.MovieCreateParams) (domain.Movie, error)
	List(ctx context.Context, movieParams params.MovieListParams) ([]ports.MovieListItem, error)
	Update(ctx context.Context, id int64, version int64, movieParams params.MovieUpdateParams) (domain.Movie, error)
	UpdatePoster(ctx context.Context, id int64, oldFilename *string, newFilename string) error
}
