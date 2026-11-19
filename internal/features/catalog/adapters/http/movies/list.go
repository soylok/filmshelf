package movies

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/helpers"
	"github.com/soylok/filmshelf/internal/features/catalog/params"
)

func (h *MovieHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	movieParams, err := parseListParams(r)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("parse query params: %w", err))
		return
	}

	movies, err := h.service.List(ctx, movieParams)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("list movies: %w", err))
		return
	}
	moviesDTO := dto.MovieListItemsToDTO(movies)
	responder.JSONResponse(moviesDTO, http.StatusOK)

}

func parseListParams(r *http.Request) (params.MovieListParams, error) {
	var movieParams params.MovieListParams

	limit, ok, err := helpers.GetInt64QueryParam(r, "limit")
	if err != nil {
		return params.MovieListParams{}, err
	}
	if ok {
		movieParams.Limit = &limit
	}

	offset, ok, err := helpers.GetInt64QueryParam(r, "offset")
	if err != nil {
		return params.MovieListParams{}, err
	}
	if ok {
		movieParams.Offset = &offset
	}

	q := r.URL.Query().Get("q")
	if q != "" {
		movieParams.Q = &q
	}

	year, ok, err := helpers.GetInt32QueryParam(r, "release_year")
	if err != nil {
		return params.MovieListParams{}, err
	}
	if ok {
		movieParams.Year = &year
	}

	rating, ok, err := helpers.GetFloatQueryParam(r, "rating")
	if err != nil {
		return params.MovieListParams{}, err
	}
	if ok {
		movieParams.Rating = &rating
	}

	query := r.URL.Query()
	for _, value := range query["genre_id"] {
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return params.MovieListParams{}, &apperrors.Error{
				Type:    apperrors.ErrInvalidArgument,
				Message: fmt.Sprintf("invalid genre_id value %q", value),
				Cause:   err,
			}
		}
		movieParams.GenreIDs = append(movieParams.GenreIDs, id)
	}

	for _, value := range query["country_id"] {
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return params.MovieListParams{}, &apperrors.Error{
				Type:    apperrors.ErrInvalidArgument,
				Message: fmt.Sprintf("invalid country_id value %q", value),
				Cause:   err,
			}
		}
		movieParams.CountryIDs = append(movieParams.CountryIDs, id)
	}

	for _, value := range query["director_id"] {
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return params.MovieListParams{}, &apperrors.Error{
				Type:    apperrors.ErrInvalidArgument,
				Message: fmt.Sprintf("invalid director_id value %q", value),
				Cause:   err,
			}
		}
		movieParams.DirectorIDs = append(movieParams.DirectorIDs, id)
	}

	sort := params.MovieSort(query.Get("sort"))
	if sort != "" {
		movieParams.Sort = &sort
	}

	return movieParams, nil
}
