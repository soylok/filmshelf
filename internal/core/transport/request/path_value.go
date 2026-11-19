package request

import (
	"net/http"
	"strconv"
)

func IntPathValue(r *http.Request, key string) (int64, error) {
	value := r.PathValue(key)
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}