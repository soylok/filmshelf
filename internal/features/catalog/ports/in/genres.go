package inbound

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
)

type GenreService interface {
	List(ctx context.Context) ([]domain.Genre, error)
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, genre domain.Genre) (domain.Genre, error)
}
