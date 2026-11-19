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

func (repo *MovieRepository) Update(
	ctx context.Context,
	id int64,
	version int64,
	movieParams params.MovieUpdateParams,
) (domain.Movie, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	tx, err := repo.querier.Begin(ctx)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("begin: %w", err)
	}
	defer tx.Rollback(ctx)

	query, args := buildUpdateQuery(id, version, movieParams)

	tag, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("exec: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return domain.Movie{}, outbound.ErrConcurrentModification
	}

	if movieParams.Countries.Set {
		query := `delete from filmshelf.movie_countries
	where movie_id = $1`
		_, err = tx.Exec(ctx, query, id)
		if err != nil {
			return domain.Movie{}, fmt.Errorf("exec delete movie_countries: %w", err)
		}
		if movieParams.Countries.Value != nil && len(*movieParams.Countries.Value) > 0 {
			query := `insert into filmshelf.movie_countries(movie_id, country_id)
			values
				%s`
			args := make([]any, 0, 1+len(*movieParams.Countries.Value))
			args = append(args, id)
			for _, countryID := range *movieParams.Countries.Value {
				args = append(args, countryID)
			}
			query = buildRelationInsertQuery(query, len(*movieParams.Countries.Value))
			_, err = tx.Exec(ctx, query, args...)
			if err != nil {
				var pgxErr *pgconn.PgError
				if errors.As(err, &pgxErr) && pgxErr.Code == "23503" {
					return domain.Movie{}, &outbound.ReferenceNotFoundError{
						Entity: "country",
					}
				}
				return domain.Movie{}, fmt.Errorf("exec insert movie_countries: %w", err)
			}
		}
	}

	if movieParams.Genres.Set {
		query := `delete from filmshelf.movie_genres
	where movie_id = $1`
		_, err = tx.Exec(ctx, query, id)
		if err != nil {
			return domain.Movie{}, fmt.Errorf("exec delete movie_genres: %w", err)
		}
		if movieParams.Genres.Value != nil && len(*movieParams.Genres.Value) > 0 {
			query := `insert into filmshelf.movie_genres(movie_id, genre_id)
			values
				%s`
			args := make([]any, 0, 1+len(*movieParams.Genres.Value))
			args = append(args, id)
			for _, genreID := range *movieParams.Genres.Value {
				args = append(args, genreID)
			}
			query = buildRelationInsertQuery(query, len(*movieParams.Genres.Value))
			_, err = tx.Exec(ctx, query, args...)
			if err != nil {
				var pgxErr *pgconn.PgError
				if errors.As(err, &pgxErr) && pgxErr.Code == "23503" {
					return domain.Movie{}, &outbound.ReferenceNotFoundError{
						Entity: "genre",
					}
				}
				return domain.Movie{}, fmt.Errorf("exec insert movie_genres: %w", err)
			}
		}
	}

	if movieParams.Directors.Set {
		query := `delete from filmshelf.movie_directors
	where movie_id = $1`
		_, err = tx.Exec(ctx, query, id)
		if err != nil {
			return domain.Movie{}, fmt.Errorf("exec delete movie_directors: %w", err)
		}
		if movieParams.Directors.Value != nil && len(*movieParams.Directors.Value) > 0 {
			query := `insert into filmshelf.movie_directors(movie_id, director_id)
			values
				%s`
			args := make([]any, 0, 1+len(*movieParams.Directors.Value))
			args = append(args, id)
			for _, directorID := range *movieParams.Directors.Value {
				args = append(args, directorID)
			}
			query = buildRelationInsertQuery(query, len(*movieParams.Directors.Value))
			_, err = tx.Exec(ctx, query, args...)
			if err != nil {
				var pgxErr *pgconn.PgError
				if errors.As(err, &pgxErr) && pgxErr.Code == "23503" {
					return domain.Movie{}, &outbound.ReferenceNotFoundError{
						Entity: "director",
					}
				}
				return domain.Movie{}, fmt.Errorf("exec insert movie_directors: %w", err)
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return domain.Movie{}, err
	}
	return repo.Get(ctx, id)
}
