package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/features/auth/params"
	outbound "github.com/soylok/filmshelf/internal/features/auth/ports/out"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, loginParams params.LoginParams) (params.LoginResult, error) {
	loginParams.Email = strings.TrimSpace(loginParams.Email)
	if err := loginParams.Validate(); err != nil {
		return params.LoginResult{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}

	user, passwordHash, err := s.repo.Login(ctx, loginParams.Email)
	if err != nil {
		if errors.Is(err, outbound.ErrNotFound) {
			return params.LoginResult{}, &apperrors.Error{
				Type:    apperrors.ErrInvalidCredentials,
				Message: "invalid email or password",
				Cause:   err,
			}
		}
		return params.LoginResult{}, fmt.Errorf("get user for login: %w", err)
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(loginParams.Password),
	); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return params.LoginResult{}, &apperrors.Error{
				Type:    apperrors.ErrInvalidCredentials,
				Message: "invalid email or password",
				Cause:   err,
			}
		}
		return params.LoginResult{}, fmt.Errorf("compare hash and password: %w", err)
	}

	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return params.LoginResult{}, fmt.Errorf("generate session token: %w", err)
	}
	token := base64.RawURLEncoding.EncodeToString(tokenBytes)
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	expiresAt := time.Now().Add(s.config.SessionTTL)
	if err := s.repo.CreateSession(ctx, tokenHash, user.ID, expiresAt); err != nil {
		return params.LoginResult{}, fmt.Errorf("create session: %w", err)
	}
	return params.LoginResult{
		User:      user,
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}
