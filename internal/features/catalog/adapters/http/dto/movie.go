package dto

import (
	"github.com/soylok/filmshelf/internal/core/domain"
	transporttypes "github.com/soylok/filmshelf/internal/core/transport/types"
	"github.com/soylok/filmshelf/internal/features/catalog/ports"
)

type MovieDTO struct {
	ID                 int64         `json:"id"`
	Title              string        `json:"title"`
	Description        *string       `json:"description"`
	ReleaseYear        *int32        `json:"release_year"`
	Rating             *float64      `json:"rating"`
	Budget             *int64        `json:"budget"`
	WorldWideBoxOffice *int64        `json:"worldwide_box_office"`
	DurationMinutes    *int32        `json:"duration_minutes"`
	PosterFilename     *string       `json:"poster_filename"`
	Countries          []CountryDTO  `json:"countries"`
	Directors          []DirectorDTO `json:"directors"`
	Genres             []GenreDTO    `json:"genres"`
	Version            int64         `json:"version"`
}

type MovieListItemDTO struct {
	ID             int64    `json:"id"`
	Title          string   `json:"title"`
	ReleaseYear    *int32   `json:"release_year"`
	Rating         *float64 `json:"rating"`
	Genre          *string  `json:"genre"`
	PosterFilename *string  `json:"poster_filename"`
}

type MovieCreateRequest struct {
	Title              string   `json:"title"`
	Description        *string  `json:"description"`
	ReleaseYear        *int32   `json:"release_year"`
	Rating             *float64 `json:"rating"`
	Budget             *int64   `json:"budget"`
	WorldWideBoxOffice *int64   `json:"worldwide_box_office"`
	DurationMinutes    *int32   `json:"duration_minutes"`
	Countries          []int64  `json:"countries"`
	Directors          []int64  `json:"directors"`
	Genres             []int64  `json:"genres"`
}

type MovieUpdateRequest struct {
	Title              transporttypes.Nullable[string]  `json:"title"`
	Description        transporttypes.Nullable[string]  `json:"description"`
	ReleaseYear        transporttypes.Nullable[int32]   `json:"release_year"`
	Rating             transporttypes.Nullable[float64] `json:"rating"`
	Budget             transporttypes.Nullable[int64]   `json:"budget"`
	WorldWideBoxOffice transporttypes.Nullable[int64]   `json:"worldwide_box_office"`
	DurationMinutes    transporttypes.Nullable[int32]   `json:"duration_minutes"`
	Countries          transporttypes.Nullable[[]int64] `json:"countries"`
	Directors          transporttypes.Nullable[[]int64] `json:"directors"`
	Genres             transporttypes.Nullable[[]int64] `json:"genres"`
}

func MovieToDTO(movie domain.Movie) MovieDTO {
	return MovieDTO{
		ID:                 movie.ID,
		Title:              movie.Title,
		Description:        movie.Description,
		ReleaseYear:        movie.ReleaseYear,
		Rating:             movie.Rating,
		Budget:             movie.Budget,
		WorldWideBoxOffice: movie.WorldWideBoxOffice,
		DurationMinutes:    movie.DurationMinutes,
		PosterFilename:     movie.PosterFilename,
		Countries:          CountriesToDTO(movie.Countries),
		Directors:          DirectorsToDTO(movie.Directors),
		Genres:             GenresToDTO(movie.Genres),
		Version:            movie.Version,
	}
}

func MovieListItemToDTO(movie ports.MovieListItem) MovieListItemDTO {
	return MovieListItemDTO{
		ID:             movie.ID,
		Title:          movie.Title,
		ReleaseYear:    movie.ReleaseYear,
		Rating:         movie.Rating,
		Genre:          movie.Genre,
		PosterFilename: movie.PosterFilename,
	}
}

func MovieListItemsToDTO(movies []ports.MovieListItem) []MovieListItemDTO {
	moviesDTO := make([]MovieListItemDTO, 0, len(movies))
	for _, movie := range movies {
		moviesDTO = append(moviesDTO, MovieListItemToDTO(movie))
	}
	return moviesDTO
}
