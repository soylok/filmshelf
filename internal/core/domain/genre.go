package domain

import (
	"errors"
	"unicode/utf8"
)

type Genre struct {
	ID      int64
	Name    string
	Version int64
}

func NewGenre(name string) Genre {
	return Genre{
		ID:      UninitializedID,
		Name:    name,
		Version: UninitializedVersion,
	}
}

func (g Genre) Validate() error {
	switch nameLen := utf8.RuneCountInString(g.Name); {
	case nameLen < 1:
		return errors.New("name must not be empty")
	case nameLen > 100:
		return errors.New("name must not exceed 100 chars")
	}
	return nil
}
