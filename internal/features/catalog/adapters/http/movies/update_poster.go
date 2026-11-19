package movies

import (
	"errors"
	"fmt"
	"mime"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/soylok/filmshelf/internal/features/catalog/adapters/http/helpers"
)

const (
	posterKey         = "poster"
	maxBodySize int64 = 10 << 20
)

func (h *MovieHandler) UpdatePoster(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)
	responder := response.NewResponder(logger, w)

	id, err := helpers.GetID(r, "id")
	if err != nil {
		responder.ErrorResponse(fmt.Errorf("parse movie id: %w", err))
		return
	}

	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		responder.ErrorResponse(&apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: "invalid Content-Type header",
			Cause:   err,
		})
		return
	}
	if mediaType != "multipart/form-data" {
		responder.ErrorResponse(&apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: "Content-type must be multipart/form-data",
		})
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	file, fileHeader, err := r.FormFile(posterKey)
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			responder.ErrorResponse(&apperrors.Error{
				Type:    apperrors.ErrInvalidArgument,
				Message: "could not read poster file from form",
				Cause:   fmt.Errorf("file is required in form field 'poster': %w", err),
			})
			return
		}

		var httpErr *http.MaxBytesError
		if errors.As(err, &httpErr) {
			responder.ErrorResponse(&apperrors.Error{
				Type:    apperrors.ErrPayloadTooLarge,
				Message: "file too large",
				Cause:   err,
			})
			return
		}
		responder.ErrorResponse(err)
		return
	}
	defer func() {
		file.Close()
		r.MultipartForm.RemoveAll()
	}()

	replaced, err := h.service.UpdatePoster(ctx, id, file, fileHeader.Filename, fileHeader.Size)
	if err != nil {
		responder.ErrorResponse(err)
		return
	}
	if !replaced {
		responder.CreatedResponse(r.URL.Path)
		return
	}
	responder.NoContentResponse()
}
