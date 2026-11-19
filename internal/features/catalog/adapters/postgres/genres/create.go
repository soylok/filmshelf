package genres

import (
	"context"
	"errors"
	"fmt"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/models"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *GenreRepository) Create(ctx context.Context, genre domain.Genre) (domain.Genre, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `insert into filmshelf.genres(name)
	values
		($1)
	returning id, name, version`
	var genreModel models.GenreModel
	row := repo.querier.QueryRow(ctx, query, genre.Name)
	err := row.Scan(
		&genreModel.ID,
		&genreModel.Name,
		&genreModel.Version,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.Genre{}, outbound.ErrAlreadyExists
		}
		return domain.Genre{}, fmt.Errorf("scan genre: %w", err)
	}
	return genreModel.ToDomain(), nil
}
