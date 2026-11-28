package movies

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/helpers"
	"github.com/soylok/filmshelf/internal/core/transport/response"
)

func (h *MovieHandler) GetPoster(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	id, err := helpers.GetID(r, "id")
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("parse movie id: %w", err))
		return
	}

	filename, err := h.service.GetPosterFilename(ctx, id)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("get poster filename: %w", err))
		return
	}
	posterPath := filepath.Join(h.postersDir, *filename)
	http.ServeFile(w, r, posterPath)
}
