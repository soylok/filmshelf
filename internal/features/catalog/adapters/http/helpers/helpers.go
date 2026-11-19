package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/transport/request"
)

func DecodeJSON(r io.Reader, dst any) error {
	if err := json.NewDecoder(r).Decode(dst); err != nil {
		return &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: "invalid JSON body",
			Cause:   err,
		}
	}
	return nil
}

func GetID(r *http.Request, key string) (int64, error) {
	id, err := request.IntPathValue(r, key)
	if err != nil {
		return 0, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: fmt.Sprintf("invalid id value %q", r.PathValue(key)),
			Cause:   err,
		}
	}
	return id, nil
}

func GetInt64QueryParam(r *http.Request, key string) (int64, bool, error) {
	value, ok, err := request.Int64QueryParam(r, key)
	if err != nil {
		return 0, true, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: fmt.Sprintf("invalid %s value %q", key, r.URL.Query().Get(key)),
			Cause:   err,
		}
	}
	return value, ok, nil
}

func GetInt32QueryParam(r *http.Request, key string) (int32, bool, error) {
	value, ok, err := request.Int32QueryParam(r, key)
	if err != nil {
		return 0, true, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: fmt.Sprintf("invalid %s value %q", key, r.URL.Query().Get(key)),
			Cause:   err,
		}
	}
	return value, ok, nil
}

func GetFloatQueryParam(r *http.Request, key string) (float64, bool, error) {
	value, ok, err := request.FloatQueryParam(r, key)
	if err != nil {
		return 0, true, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: fmt.Sprintf("invalid %s value %q", key, r.URL.Query().Get(key)),
			Cause:   err,
		}
	}
	return value, ok, nil
}
