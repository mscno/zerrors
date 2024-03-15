package zerrors

import (
	"fmt"
)

var (
	_ DeadlineExceeded = (*DeadlineExceededError)(nil)
	_ Error            = (*DeadlineExceededError)(nil)
)

type DeadlineExceeded interface {
	error
	IsDeadlineExceeded()
}

type DeadlineExceededError struct {
	*Zerror
}

func ThrowDeadlineExceeded(parent error, id, message string) error {
	return &DeadlineExceededError{CreateZerror(parent, id, message)}
}

func ThrowDeadlineExceededf(parent error, id, format string, a ...interface{}) error {
	return ThrowDeadlineExceeded(parent, id, fmt.Sprintf(format, a...))
}

func (err *DeadlineExceededError) IsDeadlineExceeded() {}

func IsDeadlineExceeded(err error) bool {
	_, ok := err.(DeadlineExceeded)
	return ok
}

func (err *DeadlineExceededError) Is(target error) bool {
	t, ok := target.(*DeadlineExceededError)
	if !ok {
		return false
	}
	return err.Zerror.Is(t.Zerror)
}

func (err *DeadlineExceededError) Unwrap() error {
	return err.Zerror
}
