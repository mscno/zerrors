package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Internal = (*InternalError)(nil)
	_ Error    = (*InternalError)(nil)
)

type Internal interface {
	error
	IsInternal()
}

type InternalError struct {
	*Zerror
}

func ThrowInternal(parent error, id string, message string) error {
	return &InternalError{CreateZerror(parent, id, message)}
}

func ThrowInternalf(parent error, id string, format string, a ...interface{}) error {
	return ThrowInternal(parent, id, fmt.Sprintf(format, a...))
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
