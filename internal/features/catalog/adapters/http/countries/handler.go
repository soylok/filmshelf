package countries

import (
	"net/http"

	"github.com/soylok/filmshelf/internal/core/transport/server"
	inbound "github.com/soylok/filmshelf/internal/features/catalog/ports/in"
)

type CountryHandler struct {
	service inbound.CountryService
}

func NewCountryHandler(service inbound.CountryService) *CountryHandler {
	return &CountryHandler{
		service: service,
	}
}

func (h *CountryHandler) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodGet, Endpoint: "/countries", Handler: h.List},
	}
}
