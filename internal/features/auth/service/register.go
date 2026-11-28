package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/domain"
	"github.com/soylok/filmshelf/internal/features/auth/params"
	outbound "github.com/soylok/filmshelf/internal/features/auth/ports/out"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Register(ctx context.Context, regParams params.RegisterParams) (domain.User, error) {
	regParams.FirstName = strings.TrimSpace(regParams.FirstName)
	regParams.SecondName = strings.TrimSpace(regParams.SecondName)
	regParams.Email = strings.TrimSpace(regParams.Email)

	user := domain.User{
		ID:         domain.UninitializedID,
		FirstName:  regParams.FirstName,
		SecondName: regParams.SecondName,
		Email:      regParams.Email,
		Role:       domain.RoleUser,
		Version:    domain.UninitializedVersion,
	}
	if err := user.Validate(); err != nil {
		return domain.User{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}
	if err := regParams.Validate(); err != nil {
		return domain.User{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(regParams.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("generate hash: %w", err)
	}
	user.PasswordHash = string(passwordHash)

	createdUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, outbound.ErrAlreadyExists) {
			return domain.User{}, &apperrors.Error{
				Type:    apperrors.ErrConflict,
				Message: "user with this email already exists",
				Cause:   err,
			}
		}
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}
	return createdUser, nil
}
