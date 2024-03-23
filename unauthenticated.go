package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Unauthenticated = (*UnauthenticatedError)(nil)
	_ Error           = (*UnauthenticatedError)(nil)
)

const UnauthenticatedId = "Unauthenticated"

type Unauthenticated interface {
	error
	IsUnauthenticated()
}

type UnauthenticatedError struct {
	*Zerror
}

func ThrowUnauthenticated(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return &UnauthenticatedError{CreateZerror(nil, UnauthenticatedId, message)}
}

func ThrowUnauthenticatedr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &UnauthenticatedError{CreateZerror(nil, UnauthenticatedId, message)}
}

func ToUnauthenticated(parent error, id string, message string) error {
	return &UnauthenticatedError{CreateZerror(parent, id, message)}
}

func ToUnauthenticatedf(parent error, id string, format string, a ...interface{}) error {
	return ToUnauthenticated(parent, id, fmt.Sprintf(format, a...))
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
