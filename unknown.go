package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Unknown = (*UnknownError)(nil)
	_ Error   = (*UnknownError)(nil)
)

type Unknown interface {
	error
	IsUnknown()
}

type UnknownError struct {
	*Zerror
}

func ThrowUnknown(parent error, id string, message string) error {
	return &UnknownError{CreateZerror(parent, id, message)}
}

func ThrowUnknownf(parent error, id string, format string, a ...interface{}) error {
	return ThrowUnknown(parent, id, fmt.Sprintf(format, a...))
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
