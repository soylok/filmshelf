package movies

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/helpers"
	"github.com/soylok/filmshelf/internal/core/transport/response"
)

func (h *MovieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	id, err := helpers.GetID(r, "id")
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("parse movie id: %w", err))
		return
	}

	err = h.service.Delete(ctx, id)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("delete movie: %w", err))
		return
	}
	responder.NoContentResponse()
}
