package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := uuid.NewString()
			r.Header.Set("X-Request-ID", id)
			w.Header().Set("X-Request-ID", id)

			next.ServeHTTP(w, r)
		})
	}
}

func InjectLogger(appLogger *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := appLogger.WithContext(r.Context())
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func Recovery() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			appLogger := logger.FromContext(r.Context())
			defer func() {
				if err := recover(); err != nil {
					responder := response.NewResponder(appLogger, w)
					responder.InternalServerError()
					appLogger.Error(
						"HTTP panic recovered",
						zap.Any("panic", err),
						zap.ByteString("stack", debug.Stack()),
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			appLogger := logger.FromContext(r.Context())
			rw := response.NewWriter(w)

			start := time.Now()

			next.ServeHTTP(rw, r)

			msg := fmt.Sprintf(
				"%s %s %d %s",
				r.Method,
				r.URL.Path,
				rw.StatusCode(),
				time.Since(start),
			)
			appLogger.Info(
				msg,
				zap.String("request-id", r.Header.Get("X-Request-ID")),
			)
		})
	}
}
