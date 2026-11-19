package movies

import (
	"context"
	"errors"
	"fmt"

	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
	"github.com/jackc/pgx/v5"
)

func (repo *MovieRepository) DeletePoster(ctx context.Context, id int64) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `update filmshelf.movies
	set
		poster_filename = null
	where id = $1
	returning old.poster_filename`

	var filename *string
	if err := repo.querier.QueryRow(ctx, query, id).Scan(&filename); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, outbound.ErrNotFound
		}
		return nil, fmt.Errorf("delete movie: %w", err)
	}
	return filename, nil
}
