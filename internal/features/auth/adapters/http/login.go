package http

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/helpers"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/auth/params"
)

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	var loginRequest LoginUserRequest
	if err := helpers.DecodeJSON(r.Body, &loginRequest); err != nil {
		responder.ErrorResponse(fmt.Errorf("decode user login request: %w", err))
		return
	}
	loginParams := params.LoginParams{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
	result, err := h.service.Login(ctx, loginParams)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("login user: %w", err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    result.Token,
		Path:     "/",
		Expires:  result.ExpiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	userDTO := UserToDTO(result.User)
	responder.JSONResponse(userDTO, http.StatusOK)
}
