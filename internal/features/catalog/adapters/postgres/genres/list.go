package genres

import (
	"context"
	"fmt"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/models"
)

func (repo *GenreRepository) List(ctx context.Context) ([]domain.Genre, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `select
		id,
		name
	from filmshelf.genres
	order by name`
	rows, err := repo.querier.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	genresModel := make([]models.GenreModel, 0)
	for rows.Next() {
		var genreModel models.GenreModel
		err := rows.Scan(
			&genreModel.ID,
			&genreModel.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan genre: %w", err)
		}
		genresModel = append(genresModel, genreModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate genres: %w", err)
	}

	return models.GenresToDomain(genresModel), nil
}
