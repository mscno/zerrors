package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Internal = (*InternalError)(nil)
	_ Error    = (*InternalError)(nil)
)

const InternalId = "Internal"

type Internal interface {
	error
	IsInternal()
}

type InternalError struct {
	*Zerror
}

func ThrowInternal(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &InternalError{CreateZerror(nil, InternalId, message)}
}

func ToInternal(parent error, id string, message string) error {
	return &InternalError{CreateZerror(parent, id, message)}
}

func ToInternalf(parent error, id string, format string, a ...interface{}) error {
	return ToInternal(parent, id, fmt.Sprintf(format, a...))
}

func (err *InternalError) IsInternal() {}

func (err *InternalError) Is(target error) bool {
	return IsInternal(target)
}

func IsInternal(err error) bool {
	var possibleError *InternalError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
