package movies

import (
	"fmt"
	"strings"

	"github.com/soylok/filmshelf/internal/features/catalog/params"
)

func buildRelationInsertQuery(query string, n int) string {
	valueGroups := make([]string, 0, n)
	cnt := 2
	for range n {
		valueGroups = append(
			valueGroups,
			fmt.Sprintf("($1, $%d)", cnt),
		)
		cnt++
	}
	return fmt.Sprintf(query, strings.Join(valueGroups, ",\n"))
}

func buildListQuery(movieParams params.MovieListParams) (string, []any) {
	query := `select
    	m.id,
    	m.title,
    	m.release_year,
    	m.rating,
        m.poster_filename,
    	(
    	    select g.name
    	    from filmshelf.movie_genres mg
    	    join filmshelf.genres g on g.id = mg.genre_id
    	    where mg.movie_id = m.id
    	    order by g.name, g.id
    	    limit 1
    	) as genre
	from filmshelf.movies m
	%s
	order by %s
	limit $1 offset $2`

	subQuery := make([]string, 0)
	args := []any{*movieParams.Limit, *movieParams.Offset}
	n := 3
	if movieParams.Q != nil {
		subQuery = append(subQuery, fmt.Sprintf(`m.title ilike $%d`, n))
		replacer := strings.NewReplacer(
			`\`, `\\`,
			`%`, `\%`,
			`_`, `\_`,
		)
		pattern := "%" + replacer.Replace(*movieParams.Q) + "%"
		args = append(args, pattern)
		n++
	}

	if movieParams.Year != nil {
		subQuery = append(subQuery, fmt.Sprintf(`m.release_year = $%d`, n))
		args = append(args, *movieParams.Year)
		n++
	}

	if movieParams.Rating != nil {
		subQuery = append(subQuery, fmt.Sprintf(`m.rating = $%d`, n))
		args = append(args, *movieParams.Rating)
		n++
	}

	if len(movieParams.GenreIDs) > 0 {
		subQuery = append(subQuery, fmt.Sprintf(`array(
			select
				genre_id
			from filmshelf.movie_genres mg
			where mg.movie_id = m.id
		) @> $%d::bigint[]`, n))
		args = append(args, movieParams.GenreIDs)
		n++
	}

	if len(movieParams.CountryIDs) > 0 {
		subQuery = append(subQuery, fmt.Sprintf(`array(
			select
				country_id
			from filmshelf.movie_countries mc
			where mc.movie_id = m.id
		) @> $%d::bigint[]`, n))
		args = append(args, movieParams.CountryIDs)
		n++
	}

	if len(movieParams.DirectorIDs) > 0 {
		subQuery = append(subQuery, fmt.Sprintf(`array(
			select
				director_id
			from filmshelf.movie_directors md
			where md.movie_id = m.id
		) @> $%d::bigint[]`, n))
		args = append(args, movieParams.DirectorIDs)
		n++
	}

	var whereCondition string
	if len(subQuery) > 0 {
		whereCondition = `where ` + strings.Join(subQuery, "\n\tand ")
	}

	sortValue := string(*movieParams.Sort)
	separator := strings.LastIndexByte(sortValue, '_')
	field, order := sortValue[:separator], sortValue[separator+1:]
	sortCondition := fmt.Sprintf("%s %s nulls last", field, order)
	if field != "id" {
		sortCondition += ", id asc"
	}
	query = fmt.Sprintf(query, whereCondition, sortCondition)
	return query, args
}

func buildUpdateQuery(
	id, version int64,
	movieParams params.MovieUpdateParams,
) (string, []any) {
	query := `update filmshelf.movies
		set
			%s
		where id = $1 and version = $2`
	valueGroups := []string{"version = version + 1"}
	args := []any{id, version}
	n := 3
	add := func(field string, value any) {
		args = append(args, value)
		valueGroups = append(valueGroups, fmt.Sprintf("%s = $%d", field, n))
		n++
	}

	if movieParams.Title.Set {
		add("title", movieParams.Title.Value)
	}
	if movieParams.Description.Set {
		add("description", movieParams.Description.Value)
	}
	if movieParams.ReleaseYear.Set {
		add("release_year", movieParams.ReleaseYear.Value)
	}
	if movieParams.Rating.Set {
		add("rating", movieParams.Rating.Value)
	}
	if movieParams.Budget.Set {
		add("budget", movieParams.Budget.Value)
	}
	if movieParams.WorldWideBoxOffice.Set {
		add("worldwide_box_office", movieParams.WorldWideBoxOffice.Value)
	}
	if movieParams.DurationMinutes.Set {
		add("duration_minutes", movieParams.DurationMinutes.Value)
	}

	setCondition := strings.Join(valueGroups, ",\n")
	query = fmt.Sprintf(query, setCondition)
	return query, args
}
