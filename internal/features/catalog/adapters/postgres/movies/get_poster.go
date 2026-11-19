package movies

import (
	"context"
	"errors"
	"fmt"

	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
	"github.com/jackc/pgx/v5"
)

func (repo *MovieRepository) GetPosterFilename(ctx context.Context, movieID int64) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `select
		poster_filename
	from filmshelf.movies
	where id = $1`
	row := repo.querier.QueryRow(ctx, query, movieID)
	var filename *string
	if err := row.Scan(&filename); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, outbound.ErrNotFound
		}
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return filename, nil
}
