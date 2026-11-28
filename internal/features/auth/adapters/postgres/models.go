package postgres

import "github.com/soylok/filmshelf/internal/core/domain"

type UserModel struct {
	ID            int64  `db:"id"`
	FirstName     string `db:"first_name"`
	SecondName    string `db:"second_name"`
	Email         string `db:"email"`
	EmailVerified bool   `db:"email_verified"`
	Role          string `db:"role"`
	Version       int64  `db:"version"`
}

func (u UserModel) ToDomain() domain.User {
	return domain.User{
		ID:            u.ID,
		FirstName:     u.FirstName,
		SecondName:    u.SecondName,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		Role:          domain.Role(u.Role),
		Version:       u.Version,
	}
}
