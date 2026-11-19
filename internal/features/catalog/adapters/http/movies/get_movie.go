package movies

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/helpers"
)

func (h *MovieHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	id, err := helpers.GetID(r, "id")
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("parse movie id: %w", err))
		return
	}

	movie, err := h.service.Get(ctx, id)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("get movie: %w", err))
		return
	}
	movieDTO := dto.MovieToDTO(movie)
	responder.JSONResponse(movieDTO, http.StatusOK)
}
