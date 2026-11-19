package genres

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
)

func (h *GenreHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	genres, err := h.service.List(ctx)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("list genres: %w", err))
		return
	}
	genresDTO := dto.GenresToDTO(genres)
	responder.JSONResponse(genresDTO, http.StatusOK)
}
