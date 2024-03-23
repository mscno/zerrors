package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ InvalidArgument = (*InvalidArgumentError)(nil)
	_ Error           = (*InvalidArgumentError)(nil)
)

const InvalidArgumentId = "InvalidArgument"

type InvalidArgument interface {
	error
	IsInvalidArgument()
}

type InvalidArgumentError struct {
	*Zerror
}

func ThrowInvalidArgument(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &InvalidArgumentError{CreateZerror(nil, InvalidArgumentId, message)}
}

func ToInvalidArgument(parent error, id string, message string) error {
	return &InvalidArgumentError{CreateZerror(parent, id, message)}
}

func ToInvalidArgumentf(parent error, id string, format string, a ...interface{}) error {
	return ToInvalidArgument(parent, id, fmt.Sprintf(format, a...))
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
