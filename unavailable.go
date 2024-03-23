package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Unavailable = (*UnavailableError)(nil)
	_ Error       = (*UnavailableError)(nil)
)

const UnavailableId = "Unavailable"

type Unavailable interface {
	error
	IsUnavailable()
}

type UnavailableError struct {
	*Zerror
}

func ThrowUnavailable(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return &UnavailableError{CreateZerror(nil, UnavailableId, message)}
}

func ThrowUnavailabler(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &UnavailableError{CreateZerror(nil, UnavailableId, message)}
}

func ToUnavailable(parent error, id string, message string) error {
	return &UnavailableError{CreateZerror(parent, id, message)}
}

func ToUnavailablef(parent error, id string, format string, a ...interface{}) error {
	return ToUnavailable(parent, id, fmt.Sprintf(format, a...))
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
