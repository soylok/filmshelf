package request

import (
	"net/http"
	"strconv"
)

func Int64QueryParam(r *http.Request, key string) (int64, bool, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, false, nil
	}
	num, err := strconv.ParseInt(value, 10, 64)
	return num, true, err
}

func Int32QueryParam(r *http.Request, key string) (int32, bool, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, false, nil
	}

	num, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, true, err
	}
	return int32(num), true, nil
}

func FloatQueryParam(r *http.Request, key string) (float64, bool, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, false, nil
	}

	num, err := strconv.ParseFloat(value, 64)
	return num, true, err
}
