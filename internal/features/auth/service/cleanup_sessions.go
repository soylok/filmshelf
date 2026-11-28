package auth

import (
	"context"
	"time"

	"github.com/soylok/filmshelf/internal/core/logger"
	"go.uber.org/zap"
)

func (s *AuthService) RunSessionCleanup(ctx context.Context) {
	log := logger.FromContext(ctx)

	ticker := time.NewTicker(s.config.SessionCleanupInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := s.repo.DeleteExpiredSessions(ctx); err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Error("delete expired sessions", zap.Error(err))
			}
		}
	}
}
