package http

import (
	"net/http"

	"github.com/soylok/filmshelf/internal/core/transport/server"
	inbound "github.com/soylok/filmshelf/internal/features/auth/ports/in"
)

type AuthHandler struct {
	service inbound.AuthService
}

func NewAuthHandler(service inbound.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodPost, Endpoint: "/register", Handler: h.Register},
		{Method: http.MethodPost, Endpoint: "/login", Handler: h.Login},
	}
}
