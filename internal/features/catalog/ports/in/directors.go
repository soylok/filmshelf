package inbound

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
)

type DirectorService interface {
	List(ctx context.Context) ([]domain.Director, error)
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, director domain.Director) (domain.Director, error)
}
