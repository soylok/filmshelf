package movies

import (
	"context"
	"errors"
	"fmt"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/params"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *MovieRepository) Create(ctx context.Context, movieParams params.MovieCreateParams) (domain.Movie, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	tx, err := repo.querier.Begin(ctx)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("begin: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `insert into filmshelf.movies(
		title,
		description,
		release_year,
		rating,
		budget,
		worldwide_box_office,
		duration_minutes
	) values
		($1, $2, $3, $4, $5, $6, $7)
	returning id`
	var movieID int64
	row := tx.QueryRow(
		ctx,
		query,
		movieParams.Title,
		movieParams.Description,
		movieParams.ReleaseYear,
		movieParams.Rating,
		movieParams.Budget,
		movieParams.WorldWideBoxOffice,
		movieParams.DurationMinutes,
	)
	err = row.Scan(&movieID)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("scan movie: %w", err)
	}

	if len(movieParams.Countries) > 0 {
		query = `insert into filmshelf.movie_countries(movie_id, country_id)
	values
		%s`
		query = buildRelationInsertQuery(query, len(movieParams.Countries))
		args := make([]any, 0, 1+len(movieParams.Countries))
		args = append(args, movieID)
		for _, countryID := range movieParams.Countries {
			args = append(args, countryID)
		}
		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			var pgxErr *pgconn.PgError
			if errors.As(err, &pgxErr) && pgxErr.Code == "23503" {
				return domain.Movie{}, &outbound.ReferenceNotFoundError{
					Entity: "country",
				}
			}
			return domain.Movie{}, fmt.Errorf("exec countries: %w", err)
		}
	}

	if len(movieParams.Genres) > 0 {
		query = `insert into filmshelf.movie_genres(movie_id, genre_id)
	values
		%s`
		query = buildRelationInsertQuery(query, len(movieParams.Genres))
		args := make([]any, 0, 1+len(movieParams.Genres))
		args = append(args, movieID)
		for _, genreID := range movieParams.Genres {
			args = append(args, genreID)
		}
		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			var pgxErr *pgconn.PgError
			if errors.As(err, &pgxErr) && pgxErr.Code == "23503" {
				return domain.Movie{}, &outbound.ReferenceNotFoundError{
					Entity: "genre",
				}
			}
			return domain.Movie{}, fmt.Errorf("exec genres: %w", err)
		}
	}

	if len(movieParams.Directors) > 0 {
		query = `insert into filmshelf.movie_directors(movie_id, director_id)
	values
		%s`
		query = buildRelationInsertQuery(query, len(movieParams.Directors))
		args := make([]any, 0, 1+len(movieParams.Directors))
		args = append(args, movieID)
		for _, directorID := range movieParams.Directors {
			args = append(args, directorID)
		}
		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			var pgxErr *pgconn.PgError
			if errors.As(err, &pgxErr) && pgxErr.Code == "23503" {
				return domain.Movie{}, &outbound.ReferenceNotFoundError{
					Entity: "director",
				}
			}
			return domain.Movie{}, fmt.Errorf("exec directors: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return domain.Movie{}, fmt.Errorf("commit: %w", err)
	}
	return repo.Get(ctx, movieID)
}
