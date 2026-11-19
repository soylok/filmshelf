package directors

import (
	"context"
	"fmt"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/models"
)

func (repo *DirectorRepository) List(ctx context.Context) ([]domain.Director, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `select
		id,
		first_name,
		last_name,
		middle_name,
		version
	from filmshelf.directors
	order by first_name, last_name, middle_name`
	rows, err := repo.querier.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	directorsModel := make([]models.DirectorModel, 0)
	for rows.Next() {
		var directorModel models.DirectorModel
		err := rows.Scan(
			&directorModel.ID,
			&directorModel.FirstName,
			&directorModel.LastName,
			&directorModel.MiddleName,
			&directorModel.Version,
		)
		if err != nil {
			return nil, fmt.Errorf("scan director: %w", err)
		}
		directorsModel = append(directorsModel, directorModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate directors: %w", err)
	}
	return models.DirectorsToDomain(directorsModel), nil
}
