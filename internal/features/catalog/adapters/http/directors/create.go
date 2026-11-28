package directors

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/helpers"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
)

func (h *DirectorHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	var directorRequest dto.DirectorCreateRequest
	if err := helpers.DecodeJSON(r.Body, &directorRequest); err != nil {
		responder.ErrorResponse(fmt.Errorf("decode director creation request: %w", err))
		return
	}

	director := domain.NewDirector(
		directorRequest.FirstName,
		directorRequest.LastName,
		directorRequest.MiddleName,
	)
	director, err := h.service.Create(ctx, director)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("create director: %w", err))
		return
	}
	directorDTO := dto.DirectorToDTO(director)
	responder.JSONResponse(directorDTO, http.StatusCreated)
}
