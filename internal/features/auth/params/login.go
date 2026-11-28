package params

import (
	"errors"
	"time"

	"github.com/soylok/filmshelf/internal/core/domain"
)

type LoginResult struct {
	User      domain.User
	Token     string
	ExpiresAt time.Time
}

type LoginParams struct {
	Email    string
	Password string
}

func (p LoginParams) Validate() error {
	if p.Email == "" || p.Password == "" {
		return errors.New("invalid email or password")
	}
	return nil
}
