package outbound

import (
	"context"
	"time"

	"github.com/soylok/filmshelf/internal/core/domain"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	CreateSession(ctx context.Context, tokenHash string, userID int64, expiresAt time.Time) error
	Login(ctx context.Context, email string) (domain.User, string, error)
	DeleteExpiredSessions(ctx context.Context) error
}
