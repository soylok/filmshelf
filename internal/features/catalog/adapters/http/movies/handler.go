package movies

import (
	"net/http"

	"github.com/soylok/filmshelf/internal/core/transport/server"
	inbound "github.com/soylok/filmshelf/internal/features/catalog/ports/in"
)

type MovieHandler struct {
	service    inbound.MovieService
	postersDir string
}

func NewMovieHandler(service inbound.MovieService, postersDir string) *MovieHandler {
	return &MovieHandler{
		service:    service,
		postersDir: postersDir,
	}
}

func (h *MovieHandler) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodGet, Endpoint: "/movies", Handler: h.List},
		{Method: http.MethodGet, Endpoint: "/movies/{id}", Handler: h.Get},
		{Method: http.MethodGet, Endpoint: "/movies/{id}/poster", Handler: h.GetPoster},

		{Method: http.MethodDelete, Endpoint: "/movies/{id}", Handler: h.Delete},
		{Method: http.MethodDelete, Endpoint: "/movies/{id}/poster", Handler: h.DeletePoster},

		{Method: http.MethodPost, Endpoint: "/movies", Handler: h.Create},

		{Method: http.MethodPatch, Endpoint: "/movies/{id}", Handler: h.Update},

		{Method: http.MethodPut, Endpoint: "/movies/{id}/poster", Handler: h.UpdatePoster},
	}
}
