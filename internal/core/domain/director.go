package domain

import (
	"errors"
	"unicode/utf8"
)

type Director struct {
	ID         int64
	FirstName  string
	LastName   string
	MiddleName *string
	Version    int64
}

func NewDirector(firstName, lastName string, middleName *string) Director {
	return Director{
		ID:         UninitializedID,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
		Version:    UninitializedVersion,
	}
}

func (d Director) Validate() error {
	switch firstNameLen := utf8.RuneCountInString(d.FirstName); {
	case firstNameLen < 1:
		return errors.New("first name must not be empty")
	case firstNameLen > 100:
		return errors.New("first name length must not exceed 100 chars")
	}

	switch lastNameLen := utf8.RuneCountInString(d.LastName); {
	case lastNameLen < 1:
		return errors.New("last name must not be empty")
	case lastNameLen > 100:
		return errors.New("last name length must not exceed 100 chars")
	}

	if d.MiddleName != nil {
		switch middleNameLen := utf8.RuneCountInString(*d.MiddleName); {
		case middleNameLen < 1:
			return errors.New("middle name must not be empty")
		case middleNameLen > 100:
			return errors.New("middle name length must not exceed 100 chars")
		}
	}
	return nil
}
