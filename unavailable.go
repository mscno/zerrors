package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Unavailable = (*UnavailableError)(nil)
	_ Error       = (*UnavailableError)(nil)
)

type Unavailable interface {
	error
	IsUnavailable()
}

type UnavailableError struct {
	*Zerror
}

func ThrowUnavailable(parent error, id string, message string) error {
	return &UnavailableError{CreateZerror(parent, id, message)}
}

func ThrowUnavailablef(parent error, id string, format string, a ...interface{}) error {
	return ThrowUnavailable(parent, id, fmt.Sprintf(format, a...))
}

func (err *UnavailableError) IsUnavailable() {}

func (err *UnavailableError) Is(target error) bool {
	return IsUnavailable(target)
}

func IsUnavailable(err error) bool {
	var possibleError *UnavailableError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
