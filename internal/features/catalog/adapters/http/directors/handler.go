package directors

import (
	"net/http"

	"github.com/soylok/filmshelf/internal/core/transport/server"
	inbound "github.com/soylok/filmshelf/internal/features/catalog/ports/in"
)

type DirectorHandler struct {
	service inbound.DirectorService
}

func NewDirectorHandler(service inbound.DirectorService) *DirectorHandler {
	return &DirectorHandler{
		service: service,
	}
}

func (h *DirectorHandler) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodGet, Endpoint: "/directors", Handler: h.List},

		{Method: http.MethodDelete, Endpoint: "/directors/{id}", Handler: h.Delete},

		{Method: http.MethodPost, Endpoint: "/directors", Handler: h.Create},
	}
}
