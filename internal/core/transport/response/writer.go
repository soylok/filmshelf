package response

import "net/http"

const (
	uninitializedStatusCode = -1
)

type Writer struct {
	http.ResponseWriter
	statusCode int
}

func NewWriter(w http.ResponseWriter) *Writer {
	return &Writer{
		ResponseWriter: w,
		statusCode:     uninitializedStatusCode,
	}
}

func (w *Writer) StatusCode() int {
	if w.statusCode == uninitializedStatusCode {
		return http.StatusOK
	}

	return w.statusCode
}

func (w *Writer) Write(v []byte) (int, error) {
	if w.statusCode == uninitializedStatusCode {
		w.statusCode = http.StatusOK
	}

	n, err := w.ResponseWriter.Write(v)
	return n, err
}

func (w *Writer) WriteHeader(statusCode int) {
	if w.statusCode != uninitializedStatusCode {
		return
	}

	// Exclude interim status codes.
	if statusCode > 99 && statusCode < 200 && statusCode != 101 {
		w.ResponseWriter.WriteHeader(statusCode)
		return
	}

	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}
