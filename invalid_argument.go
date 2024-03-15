package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ InvalidArgument = (*InvalidArgumentError)(nil)
	_ Error           = (*InvalidArgumentError)(nil)
)

type InvalidArgument interface {
	error
	IsInvalidArgument()
}

type InvalidArgumentError struct {
	*Zerror
}

func ThrowInvalidArgument(parent error, id string, message string) error {
	return &InvalidArgumentError{CreateZerror(parent, id, message)}
}

func ThrowInvalidArgumentf(parent error, id string, format string, a ...interface{}) error {
	return ThrowInvalidArgument(parent, id, fmt.Sprintf(format, a...))
}

func (err *InvalidArgumentError) IsInvalidArgument() {}

func (err *InvalidArgumentError) Is(target error) bool {
	return IsInvalidArgument(target)
}

func IsInvalidArgument(err error) bool {
	var possibleError *InvalidArgumentError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
