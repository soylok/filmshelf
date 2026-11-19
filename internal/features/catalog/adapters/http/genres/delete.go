package genres

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/helpers"
)

func (h *GenreHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	id, err := helpers.GetID(r, "id")
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("parse genre id: %w", err))
		return
	}

	err = h.service.Delete(ctx, id)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("delete genre: %w", err))
		return
	}
	responder.NoContentResponse()
}
