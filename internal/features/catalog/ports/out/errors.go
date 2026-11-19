package outbound

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound               = errors.New("not found")
	ErrReferenceNotFound      = errors.New("reference not found")
	ErrAlreadyExists          = errors.New("already exists")
	ErrConcurrentModification = errors.New("concurrent modification")
)

type ReferenceNotFoundError struct {
	Entity string
}

func (e *ReferenceNotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", e.Entity, ErrReferenceNotFound)
}

func (e *ReferenceNotFoundError) Unwrap() error {
	return ErrReferenceNotFound
}
