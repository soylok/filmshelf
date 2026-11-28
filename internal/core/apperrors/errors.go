package apperrors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrNotFound           = errors.New("resource not found")
	ErrConflict           = errors.New("conflict")
	ErrPayloadTooLarge    = errors.New("payload too large")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Error struct {
	Type    error
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Cause)
}

func (e *Error) Unwrap() []error {
	return []error{e.Type, e.Cause}
}
