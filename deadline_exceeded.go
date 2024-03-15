package zerrors

import (
	"errors"
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

func ThrowDeadlineExceeded(parent error, id string, message string) error {
	return &DeadlineExceededError{CreateZerror(parent, id, message)}
}

func ThrowDeadlineExceededf(parent error, id string, format string, a ...interface{}) error {
	return ThrowDeadlineExceeded(parent, id, fmt.Sprintf(format, a...))
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
