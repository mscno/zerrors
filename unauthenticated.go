package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Unauthenticated = (*UnauthenticatedError)(nil)
	_ Error           = (*UnauthenticatedError)(nil)
)

type Unauthenticated interface {
	error
	IsUnauthenticated()
}

type UnauthenticatedError struct {
	*Zerror
}

func ThrowUnauthenticated(parent error, id string, message string) error {
	return &UnauthenticatedError{CreateZerror(parent, id, message)}
}

func ThrowUnauthenticatedf(parent error, id string, format string, a ...interface{}) error {
	return ThrowUnauthenticated(parent, id, fmt.Sprintf(format, a...))
}

func (err *UnauthenticatedError) IsUnauthenticated() {}

func (err *UnauthenticatedError) Is(target error) bool {
	return IsUnauthenticated(target)
}

func IsUnauthenticated(err error) bool {
	var possibleError *UnauthenticatedError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
