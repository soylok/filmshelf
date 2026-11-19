package directors

import (
	"context"
	"fmt"

	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

func (repo *DirectorRepository) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `delete
	from filmshelf.directors
	where id = $1`

	tag, err := repo.querier.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return outbound.ErrNotFound
	}
	return nil
}
