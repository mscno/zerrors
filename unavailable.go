package zerrors

import (
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

func ThrowUnavailable(parent error, id, message string) error {
	return &UnavailableError{CreateZerror(parent, id, message)}
}

func ThrowUnavailablef(parent error, id, format string, a ...interface{}) error {
	return ThrowUnavailable(parent, id, fmt.Sprintf(format, a...))
}

func (err *UnavailableError) IsUnavailable() {}

func IsUnavailable(err error) bool {
	_, ok := err.(Unavailable)
	return ok
}

func (err *UnavailableError) Is(target error) bool {
	t, ok := target.(*UnavailableError)
	if !ok {
		return false
	}
	return err.Zerror.Is(t.Zerror)
}

func (err *UnavailableError) Unwrap() error {
	return err.Zerror
}
