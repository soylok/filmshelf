package movies

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/helpers"
	"github.com/soylok/filmshelf/internal/features/catalog/params"
)

func (h *MovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	var movieRequest dto.MovieCreateRequest
	if err := helpers.DecodeJSON(r.Body, &movieRequest); err != nil {
		responder.ErrorResponse(fmt.Errorf("decode movie creation request: %w", err))
		return
	}

	movieParams := params.MovieCreateParams{
		Title:              movieRequest.Title,
		Description:        movieRequest.Description,
		ReleaseYear:        movieRequest.ReleaseYear,
		Rating:             movieRequest.Rating,
		Budget:             movieRequest.Budget,
		WorldWideBoxOffice: movieRequest.WorldWideBoxOffice,
		DurationMinutes:    movieRequest.DurationMinutes,
		Countries:          movieRequest.Countries,
		Directors:          movieRequest.Directors,
		Genres:             movieRequest.Genres,
	}
	movie, err := h.service.Create(ctx, movieParams)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("create movie: %w", err))
		return
	}
	movieDTO := dto.MovieToDTO(movie)
	responder.JSONResponse(movieDTO, http.StatusCreated)
}
