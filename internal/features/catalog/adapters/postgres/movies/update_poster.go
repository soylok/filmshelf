package movies

import (
	"context"
	"fmt"

	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

func (repo *MovieRepository) UpdatePoster(
	ctx context.Context,
	id int64,
	oldFilename *string,
	newFilename string,
) error {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `update filmshelf.movies
	set
		poster_filename = $1,
		version = version + 1
	where id = $2
		and poster_filename is not distinct from $3`

	tag, err := repo.querier.Exec(ctx, query, newFilename, id, oldFilename)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return outbound.ErrConcurrentModification
	}
	return nil
}
