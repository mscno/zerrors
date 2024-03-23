package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ NotFound = (*NotFoundError)(nil)
	_ Error    = (*NotFoundError)(nil)
)

const NotFoundId = "NotFound"

type NotFound interface {
	error
	IsNotFound()
}

type NotFoundError struct {
	*Zerror
}

func ThrowNotFound(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &NotFoundError{CreateZerror(nil, NotFoundId, message)}
}

func ToNotFound(parent error, id string, message string) error {
	return &NotFoundError{CreateZerror(parent, id, message)}
}

func ToNotFoundf(parent error, id string, format string, a ...interface{}) error {
	return ToNotFound(parent, id, fmt.Sprintf(format, a...))
}

func (err *NotFoundError) IsNotFound() {}

func (err *NotFoundError) Is(target error) bool {
	return IsNotFound(target)
}

func IsNotFound(err error) bool {
	var possibleError *NotFoundError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
