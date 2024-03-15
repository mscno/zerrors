package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ FailedPrecondition = (*FailedPreconditionError)(nil)
	_ Error              = (*FailedPreconditionError)(nil)
)

type FailedPrecondition interface {
	error
	IsFailedPrecondition()
}

type FailedPreconditionError struct {
	*Zerror
}

func ThrowFailedPrecondition(parent error, id string, message string) error {
	return &FailedPreconditionError{CreateZerror(parent, id, message)}
}

func ThrowFailedPreconditionf(parent error, id string, format string, a ...interface{}) error {
	return ThrowFailedPrecondition(parent, id, fmt.Sprintf(format, a...))
}

func (err *FailedPreconditionError) IsFailedPrecondition() {}

func (err *FailedPreconditionError) Is(target error) bool {
	return IsFailedPrecondition(target)
}

func IsFailedPrecondition(err error) bool {
	var possibleError *FailedPreconditionError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
