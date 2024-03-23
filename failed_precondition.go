package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ FailedPrecondition = (*FailedPreconditionError)(nil)
	_ Error              = (*FailedPreconditionError)(nil)
)

const FailedPreconditionId = "FailedPrecondition"

type FailedPrecondition interface {
	error
	IsFailedPrecondition()
}

type FailedPreconditionError struct {
	*Zerror
}

func ThrowFailedPrecondition(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &FailedPreconditionError{CreateZerror(nil, FailedPreconditionId, message)}
}

func ToFailedPrecondition(parent error, id string, message string) error {
	return &FailedPreconditionError{CreateZerror(parent, id, message)}
}

func ToFailedPreconditionf(parent error, id string, format string, a ...interface{}) error {
	return ToFailedPrecondition(parent, id, fmt.Sprintf(format, a...))
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
