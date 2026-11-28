package params

import (
	"errors"
	"unicode/utf8"
)

type RegisterParams struct {
	FirstName       string
	SecondName      string
	Email           string
	Password        string
	ConfirmPassword string
}

func (p RegisterParams) Validate() error {
	switch passwordLen := utf8.RuneCountInString(p.Password); {
	case passwordLen < 8:
		return errors.New("password must contain at least 8 chars")
	case passwordLen > 50:
		return errors.New("password must not exceed 50 chars")
	}

	for _, sym := range p.Password {
		if !(sym >= '!' && sym <= '~' || sym == ' ') {
			return errors.New("password must contain only Latin letters, digits or special chars")
		}
	}
	if p.Password != p.ConfirmPassword {
		return errors.New("passwords do not match")
	}
	return nil
}
