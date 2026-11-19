package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/logger"
	"go.uber.org/zap"
)

type Responder struct {
	log *logger.Logger
	w   http.ResponseWriter
}

func NewResponder(log *logger.Logger, w http.ResponseWriter) *Responder {
	return &Responder{
		log: log,
		w:   w,
	}
}

func (resp *Responder) JSONResponse(data any, statusCode int) {
	resp.w.Header().Set("Content-Type", "application/json; charset=utf-8")

	bytes, err := json.Marshal(data)
	if err != nil {
		resp.log.Error(
			"marshal data",
			zap.Error(err),
		)

		resp.InternalServerError()
		return
	}

	resp.w.WriteHeader(statusCode)
	_, err = resp.w.Write(bytes)
	if err != nil {
		resp.log.Error("write response", zap.Error(err))
	}
}

func (resp *Responder) NoContentResponse() {
	resp.w.WriteHeader(http.StatusNoContent)
}

func (resp *Responder) CreatedResponse(location string) {
	resp.w.Header().Set("Location", location)
	resp.w.WriteHeader(http.StatusCreated)
}

func (resp *Responder) InternalServerError() {
	resp.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.w.WriteHeader(http.StatusInternalServerError)
	if _, err := resp.w.Write([]byte(`{"error":"internal server error"}`)); err != nil {
		resp.log.Error("write response", zap.Error(err))
	}
}

func (resp *Responder) ErrorResponse(err error) {
	msg := "internal server error"
	statusCode := http.StatusInternalServerError

	var appErr *apperrors.Error
	if errors.As(err, &appErr) {
		msg = appErr.Message

		switch appErr.Type {
		case apperrors.ErrInvalidArgument:
			statusCode = http.StatusBadRequest

		case apperrors.ErrNotFound:
			statusCode = http.StatusNotFound

		case apperrors.ErrConflict:
			statusCode = http.StatusConflict

		case apperrors.ErrPayloadTooLarge:
			statusCode = http.StatusRequestEntityTooLarge
		}
	}

	if statusCode >= http.StatusInternalServerError {
		resp.log.Error(msg, zap.Error(err))
		resp.InternalServerError()
		return
	}
	resp.log.Warn(msg, zap.Error(err))

	errorMsg := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}
	resp.JSONResponse(errorMsg, statusCode)
}
