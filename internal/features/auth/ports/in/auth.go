package inbound

import (
	"context"

	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/auth/params"
)

type AuthService interface {
	Register(ctx context.Context, regParams params.RegisterParams) (domain.User, error)
	Login(ctx context.Context, loginParams params.LoginParams) (params.LoginResult, error)
}
