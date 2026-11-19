package movies

import (
	"context"
	"errors"
	"fmt"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/postgres/models"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
	"github.com/jackc/pgx/v5"
)

func (repo *MovieRepository) Get(ctx context.Context, id int64) (domain.Movie, error) {
	batch := &pgx.Batch{}
	batch.Queue(`select
		id,
		title,
		description,
		release_year,
		rating,
		budget,
		worldwide_box_office,
		duration_minutes,
		poster_filename,
		version
	from filmshelf.movies
	where id = $1`,
		id,
	)
	batch.Queue(`select
		c.id,
		c.code
	from filmshelf.movie_countries mc
	join filmshelf.countries c on c.id = mc.country_id
	where mc.movie_id = $1
	order by code`,
		id,
	)
	batch.Queue(`select
		d.id,
    	d.first_name,
    	d.last_name,
    	d.middle_name,
		d.version
	from filmshelf.movie_directors md
	join filmshelf.directors d on d.id = md.director_id
	where md.movie_id = $1
	order by last_name, first_name, middle_name`,
		id,
	)
	batch.Queue(`select
		g.id,
		g.name,
		g.version
	from filmshelf.movie_genres mg
	join filmshelf.genres g on g.id = mg.genre_id
	where mg.movie_id = $1
	order by name`,
		id,
	)

	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	br := repo.querier.SendBatch(ctx, batch)
	defer br.Close()

	var movieModel models.MovieModel
	row := br.QueryRow()
	err := row.Scan(
		&movieModel.ID,
		&movieModel.Title,
		&movieModel.Description,
		&movieModel.ReleaseYear,
		&movieModel.Rating,
		&movieModel.Budget,
		&movieModel.WorldWideBoxOffice,
		&movieModel.DurationMinutes,
		&movieModel.PosterFilename,
		&movieModel.Version,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Movie{}, outbound.ErrNotFound
		}
		return domain.Movie{}, fmt.Errorf("scan row: %w", err)
	}

	rows, err := br.Query()
	if err != nil {
		return domain.Movie{}, fmt.Errorf("query rows: %w", err)
	}
	for rows.Next() {
		var country models.CountryModel
		if err := rows.Scan(&country.ID, &country.Code); err != nil {
			return domain.Movie{}, fmt.Errorf("scan movie country: %w", err)
		}
		movieModel.Countries = append(movieModel.Countries, country)
	}
	if err := rows.Err(); err != nil {
		return domain.Movie{}, fmt.Errorf("iterate movie countries: %w", err)
	}

	rows, err = br.Query()
	if err != nil {
		return domain.Movie{}, fmt.Errorf("query rows: %w", err)
	}
	for rows.Next() {
		var director models.DirectorModel
		if err := rows.Scan(&director.ID, &director.FirstName, &director.LastName, &director.MiddleName, &director.Version); err != nil {
			return domain.Movie{}, fmt.Errorf("scan movie director: %w", err)
		}
		movieModel.Directors = append(movieModel.Directors, director)
	}
	if err := rows.Err(); err != nil {
		return domain.Movie{}, fmt.Errorf("iterate movie directors: %w", err)
	}

	rows, err = br.Query()
	if err != nil {
		return domain.Movie{}, fmt.Errorf("query rows: %w", err)
	}
	for rows.Next() {
		var genre models.GenreModel
		if err := rows.Scan(&genre.ID, &genre.Name, &genre.Version); err != nil {
			return domain.Movie{}, fmt.Errorf("scan movie genre: %w", err)
		}
		movieModel.Genres = append(movieModel.Genres, genre)
	}
	if err := rows.Err(); err != nil {
		return domain.Movie{}, fmt.Errorf("iterate movie genres: %w", err)
	}

	return movieModel.ToDomain(), nil
}
