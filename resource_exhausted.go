package zerrors

import (
	"fmt"
)

var (
	_ ResourceExhausted = (*ResourceExhaustedError)(nil)
	_ Error             = (*ResourceExhaustedError)(nil)
)

type ResourceExhausted interface {
	error
	IsResourceExhausted()
}

type ResourceExhaustedError struct {
	*Zerror
}

func ThrowResourceExhausted(parent error, id, message string) error {
	return &ResourceExhaustedError{CreateZerror(parent, id, message)}
}

func ThrowResourceExhaustedf(parent error, id, format string, a ...interface{}) error {
	return ThrowResourceExhausted(parent, id, fmt.Sprintf(format, a...))
}

func (err *ResourceExhaustedError) IsResourceExhausted() {}

func IsResourceExhausted(err error) bool {
	//nolint:errorlint
	_, ok := err.(ResourceExhausted)
	return ok
}

func (err *ResourceExhaustedError) Is(target error) bool {
	//nolint:errorlint
	t, ok := target.(*ResourceExhaustedError)
	if !ok {
		return false
	}
	return err.Zerror.Is(t.Zerror)
}

func (err *ResourceExhaustedError) Unwrap() error {
	return err.Zerror
}
