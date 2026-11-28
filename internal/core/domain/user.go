package domain

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID            int64
	FirstName     string
	SecondName    string
	Email         string
	EmailVerified bool
	PasswordHash  string
	Role          Role
	Version       int64
}

var emailPattern = regexp.MustCompile(`^[^@\s]+@[^@\s]+$`)

func (u User) Validate() error {
	switch firstNameLen := utf8.RuneCountInString(u.FirstName); {
	case firstNameLen < 1:
		return errors.New("first name must not be empty")
	case firstNameLen > 100:
		return errors.New("first name must not exceed 100 chars")
	}

	switch secondNameLen := utf8.RuneCountInString(u.SecondName); {
	case secondNameLen < 1:
		return errors.New("second name must not be empty")
	case secondNameLen > 100:
		return errors.New("second name must not exceed 100 chars")
	}

	if !emailPattern.MatchString(u.Email) {
		return errors.New("invalid email format")
	}
	return nil
}
