package genres

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/helpers"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
)

func (h *GenreHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	var genreRequest dto.GenreCreateRequest
	if err := helpers.DecodeJSON(r.Body, &genreRequest); err != nil {
		responder.ErrorResponse(fmt.Errorf("decode genre creation request: %w", err))
	}

	genre := domain.NewGenre(
		genreRequest.Name,
	)
	genre, err := h.service.Create(ctx, genre)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("create genre: %w", err))
		return
	}
	genreDTO := dto.GenreToDTO(genre)
	responder.JSONResponse(genreDTO, http.StatusCreated)
}
