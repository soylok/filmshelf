package auth

import outbound "github.com/soylok/filmshelf/internal/features/auth/ports/out"

type AuthService struct {
	repo   outbound.AuthRepository
	config Config
}

func NewAuthService(repository outbound.AuthRepository, config Config) *AuthService {
	return &AuthService{
		repo:   repository,
		config: config,
	}
}
