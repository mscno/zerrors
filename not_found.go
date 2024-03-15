package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ NotFound = (*NotFoundError)(nil)
	_ Error    = (*NotFoundError)(nil)
)

type NotFound interface {
	error
	IsNotFound()
}

type NotFoundError struct {
	*Zerror
}

func ThrowNotFound(parent error, id string, message string) error {
	return &NotFoundError{CreateZerror(parent, id, message)}
}

func ThrowNotFoundf(parent error, id string, format string, a ...interface{}) error {
	return ThrowNotFound(parent, id, fmt.Sprintf(format, a...))
}

func (err *NotFoundError) IsNotFound() {}

func (err *NotFoundError) Is(target error) bool {
	return IsNotFound(target)
}

func IsNotFound(err error) bool {
	var possibleError *NotFoundError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
