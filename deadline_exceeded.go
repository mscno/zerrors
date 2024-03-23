package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ DeadlineExceeded = (*DeadlineExceededError)(nil)
	_ Error            = (*DeadlineExceededError)(nil)
)

const DeadlineExceededId = "DeadlineExceeded"

type DeadlineExceeded interface {
	error
	IsDeadlineExceeded()
}

type DeadlineExceededError struct {
	*Zerror
}

func ThrowDeadlineExceeded(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return &DeadlineExceededError{CreateZerror(nil, DeadlineExceededId, message)}
}

func ThrowDeadlineExceededr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &DeadlineExceededError{CreateZerror(nil, DeadlineExceededId, message)}
}

func ToDeadlineExceeded(parent error, id string, message string) error {
	return &DeadlineExceededError{CreateZerror(parent, id, message)}
}

func ToDeadlineExceededf(parent error, id string, format string, a ...interface{}) error {
	return ToDeadlineExceeded(parent, id, fmt.Sprintf(format, a...))
}

func (err *DeadlineExceededError) IsDeadlineExceeded() {}

func (err *DeadlineExceededError) Is(target error) bool {
	return IsDeadlineExceeded(target)
}

func IsDeadlineExceeded(err error) bool {
	var possibleError *DeadlineExceededError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
