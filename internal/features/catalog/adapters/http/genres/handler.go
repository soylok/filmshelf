package genres

import (
	"net/http"

	"github.com/soylok/filmshelf/internal/core/transport/server"
	inbound "github.com/soylok/filmshelf/internal/features/catalog/ports/in"
)

type GenreHandler struct {
	service inbound.GenreService
}

func NewGenreHandler(service inbound.GenreService) *GenreHandler {
	return &GenreHandler{
		service: service,
	}
}

func (h *GenreHandler) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodGet, Endpoint: "/genres", Handler: h.List},

		{Method: http.MethodDelete, Endpoint: "/genres/{id}", Handler: h.Delete},

		{Method: http.MethodPost, Endpoint: "/genres", Handler: h.Create},
	}
}
