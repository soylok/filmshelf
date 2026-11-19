package directors

import (
	"context"
	"fmt"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/models"
)

func (repo *DirectorRepository) Create(ctx context.Context, director domain.Director) (domain.Director, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `insert into filmshelf.directors(first_name, last_name, middle_name)
	values
		($1, $2, $3)
	returning id, first_name, last_name, middle_name, version`
	var directorModel models.DirectorModel
	row := repo.querier.QueryRow(ctx, query, director.FirstName, director.LastName, director.MiddleName)
	err := row.Scan(
		&directorModel.ID,
		&directorModel.FirstName,
		&directorModel.LastName,
		&directorModel.MiddleName,
		&directorModel.Version,
	)
	if err != nil {
		return domain.Director{}, fmt.Errorf("scan director: %w", err)
	}
	return directorModel.ToDomain(), nil
}
