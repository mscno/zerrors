package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Unknown = (*UnknownError)(nil)
	_ Error   = (*UnknownError)(nil)
)

const UnknownId = "Unknown"

type Unknown interface {
	error
	IsUnknown()
}

type UnknownError struct {
	*Zerror
}

func ThrowUnknown(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &UnknownError{CreateZerror(nil, UnknownId, message)}
}

func ToUnknown(parent error, id string, message string) error {
	return &UnknownError{CreateZerror(parent, id, message)}
}

func ToUnknownf(parent error, id string, format string, a ...interface{}) error {
	return ToUnknown(parent, id, fmt.Sprintf(format, a...))
}

func (err *UnknownError) IsUnknown() {}

func (err *UnknownError) Is(target error) bool {
	return IsUnknown(target)
}

func IsUnknown(err error) bool {
	var possibleError *UnknownError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
