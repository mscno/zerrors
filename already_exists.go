package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ AlreadyExists = (*AlreadyExistsError)(nil)
	_ Error         = (*AlreadyExistsError)(nil)
)

type AlreadyExists interface {
	error
	IsAlreadyExists()
}

type AlreadyExistsError struct {
	*Zerror
}

func ThrowAlreadyExists(parent error, id string, message string) error {
	return &AlreadyExistsError{CreateZerror(parent, id, message)}
}

func ThrowAlreadyExistsf(parent error, id string, format string, a ...interface{}) error {
	return ThrowAlreadyExists(parent, id, fmt.Sprintf(format, a...))
}

func (err *AlreadyExistsError) IsAlreadyExists() {}

func (err *AlreadyExistsError) Is(target error) bool {
	return IsAlreadyExists(target)
}

func IsAlreadyExists(err error) bool {
	var possibleError *AlreadyExistsError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
