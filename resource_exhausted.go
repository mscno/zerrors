package zerrors

import (
	"errors"
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

func ThrowResourceExhausted(parent error, id string, message string) error {
	return &ResourceExhaustedError{CreateZerror(parent, id, message)}
}

func ThrowResourceExhaustedf(parent error, id string, format string, a ...interface{}) error {
	return ThrowResourceExhausted(parent, id, fmt.Sprintf(format, a...))
}

func (err *ResourceExhaustedError) IsResourceExhausted() {}

func (err *ResourceExhaustedError) Is(target error) bool {
	return IsResourceExhausted(target)
}

func IsResourceExhausted(err error) bool {
	var possibleError *ResourceExhaustedError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
