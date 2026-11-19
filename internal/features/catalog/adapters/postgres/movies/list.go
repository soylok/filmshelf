package movies

import (
	"context"
	"fmt"

	"github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/models"
	"github.com/soylok/filmshelf/internal/features/catalog/params"
	"github.com/soylok/filmshelf/internal/features/catalog/ports"
)

func (repo *MovieRepository) List(ctx context.Context, movieParams params.MovieListParams) ([]ports.MovieListItem, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query, args := buildListQuery(movieParams)

	movieListModel := make([]models.MovieListItemModel, 0)
	rows, err := repo.querier.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var movieItem models.MovieListItemModel
		err := rows.Scan(
			&movieItem.ID,
			&movieItem.Title,
			&movieItem.ReleaseYear,
			&movieItem.Rating,
			&movieItem.PosterFilename,
			&movieItem.Genre,
		)
		if err != nil {
			return nil, fmt.Errorf("scan movie: %w", err)
		}
		movieListModel = append(movieListModel, movieItem)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate movies: %w", err)
	}
	return models.MoviesToDomain(movieListModel), nil

}
