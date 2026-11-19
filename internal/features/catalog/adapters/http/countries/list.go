package countries

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/dto"
)

func (h *CountryHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	countries, err := h.service.List(ctx)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("list countries: %w", err))
		return
	}
	countriesDTO := dto.CountriesToDTO(countries)
	responder.JSONResponse(countriesDTO, http.StatusOK)
}
