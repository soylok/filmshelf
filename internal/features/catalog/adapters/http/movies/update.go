package movies

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/helpers"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
	"github.com/soylok/filmshelf/internal/features/catalog/params"
)

func (h *MovieHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	id, err := helpers.GetID(r, "id")
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("parse movie id: %w", err))
		return
	}
	var movieRequest dto.MovieUpdateRequest
	if err := helpers.DecodeJSON(r.Body, &movieRequest); err != nil {
		responder.ErrorResponse(fmt.Errorf("decode movie update request: %w", err))
		return
	}
	movieParams := params.MovieUpdateParams{
		Title:              movieRequest.Title.Nullable,
		Description:        movieRequest.Description.Nullable,
		ReleaseYear:        movieRequest.ReleaseYear.Nullable,
		Rating:             movieRequest.Rating.Nullable,
		Budget:             movieRequest.Budget.Nullable,
		WorldWideBoxOffice: movieRequest.WorldWideBoxOffice.Nullable,
		DurationMinutes:    movieRequest.DurationMinutes.Nullable,
		Countries:          movieRequest.Countries.Nullable,
		Directors:          movieRequest.Directors.Nullable,
		Genres:             movieRequest.Genres.Nullable,
	}
	movie, err := h.service.Update(ctx, id, movieParams)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("update movie: %w", err))
		return
	}
	movieDTO := dto.MovieToDTO(movie)
	responder.JSONResponse(movieDTO, http.StatusOK)
}
