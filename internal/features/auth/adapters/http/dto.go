package http

import "github.com/soylok/filmshelf/internal/core/domain"

type UserDTO struct {
	ID            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	SecondName    string `json:"second_name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Role          string `json:"role"`
	Version       int64  `json:"version"`
}

type RegisterUserRequest struct {
	FirstName       string `json:"first_name"`
	SecondName      string `json:"second_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserToDTO(user domain.User) UserDTO {
	return UserDTO{
		ID:            user.ID,
		FirstName:     user.FirstName,
		SecondName:    user.SecondName,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Role:          string(user.Role),
		Version:       user.Version,
	}
}
