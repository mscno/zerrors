package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ ResourceExhausted = (*ResourceExhaustedError)(nil)
	_ Error             = (*ResourceExhaustedError)(nil)
)

const ResourceExhaustedId = "ResourceExhausted"

type ResourceExhausted interface {
	error
	IsResourceExhausted()
}

type ResourceExhaustedError struct {
	*Zerror
}

func ThrowResourceExhausted(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &ResourceExhaustedError{CreateZerror(nil, ResourceExhaustedId, message)}
}

func ToResourceExhausted(parent error, id string, message string) error {
	return &ResourceExhaustedError{CreateZerror(parent, id, message)}
}

func ToResourceExhaustedf(parent error, id string, format string, a ...interface{}) error {
	return ToResourceExhausted(parent, id, fmt.Sprintf(format, a...))
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
