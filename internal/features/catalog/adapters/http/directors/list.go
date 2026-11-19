package directors

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
)

func (h *DirectorHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	directors, err := h.service.List(ctx)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("list directors: %w", err))
		return
	}
	directorsDTO := dto.DirectorsToDTO(directors)
	responder.JSONResponse(directorsDTO, http.StatusOK)
}
