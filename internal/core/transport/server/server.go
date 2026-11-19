package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	applogger "github.com/soylok/filmshelf/internal/core/logger"
	"github.com/soylok/filmshelf/internal/core/transport/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux         *http.ServeMux
	appLogger   *applogger.Logger
	middlewares []middleware.Middleware
	config      Config
}

func NewHTTPServer(
	config Config,
	logger *applogger.Logger,
	middlewares []middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:         http.NewServeMux(),
		appLogger:   logger,
		config:      config,
		middlewares: middlewares,
	}
}

func (s *HTTPServer) RegisterRouters(routers ...*Router) {
	for _, router := range routers {
		prefix := fmt.Sprintf("/api/%s", router.v)
		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, router.mux))
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	handler := middleware.Chain(s.mux, s.middlewares)

	server := &http.Server{
		Addr:        s.config.Addr,
		Handler:     handler,
		ReadTimeout: s.config.ReadTimeout,
		IdleTimeout: s.config.IdleTimeout,
	}

	chErr := make(chan error, 2)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			chErr <- err
			return
		}
	}()

	select {
	case <-ctx.Done():
		s.appLogger.Info("HTTP server shutting down")

		ctxSd, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer cancel()
		start := time.Now()
		if err := server.Shutdown(ctxSd); err != nil {
			s.appLogger.Warn("HTTP server shutdown failed")
			if errClose := server.Close(); errClose != nil {
				s.appLogger.Warn("HTTP server force close failed", zap.Error(errClose))
				return errors.Join(err, errClose)
			}
			s.appLogger.Warn("HTTP server closing forcefully", zap.Error(err))
			return err
		}
		s.appLogger.Info(
			fmt.Sprintf("HTTP server shutdown successfully completed in %s", time.Since(start)),
		)
		return nil
	case err := <-chErr:
		return err
	}
}
