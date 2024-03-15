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

func ThrowAlreadyExists(parent error, id, message string) error {
	return &AlreadyExistsError{CreateZerror(parent, id, message)}
}

func ThrowAlreadyExistsf(parent error, id, format string, a ...interface{}) error {
	return &AlreadyExistsError{CreateZerror(parent, id, fmt.Sprintf(format, a...))}
}

func (err *AlreadyExistsError) IsAlreadyExists() {}

func (err *AlreadyExistsError) Is(target error) bool {
	var possibleError *AlreadyExistsError
	if errors.As(target, &possibleError) {
		target = possibleError
		return true
	}
	return false
}

func IsErrorAlreadyExists(err error) bool {
	var possibleError *AlreadyExistsError
	if errors.As(err, &possibleError) {
		err = possibleError
		return true
	}
	return false
}

func (err *AlreadyExistsError) Unwrap() error {
	return err.Zerror
}
