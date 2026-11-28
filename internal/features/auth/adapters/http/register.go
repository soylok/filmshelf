package http

import (
	"fmt"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/helpers"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/auth/params"
)

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	var registerRequest RegisterUserRequest
	if err := helpers.DecodeJSON(r.Body, &registerRequest); err != nil {
		responder.ErrorResponse(fmt.Errorf("decode user register request: %w", err))
		return
	}

	registerParams := params.RegisterParams{
		FirstName:       registerRequest.FirstName,
		SecondName:      registerRequest.SecondName,
		Email:           registerRequest.Email,
		Password:        registerRequest.Password,
		ConfirmPassword: registerRequest.ConfirmPassword,
	}
	user, err := h.service.Register(ctx, registerParams)
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("register user: %w", err))
		return
	}
	userDTO := UserToDTO(user)
	responder.JSONResponse(userDTO, http.StatusCreated)
}
