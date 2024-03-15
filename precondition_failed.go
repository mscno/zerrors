package zerrors

import (
	"fmt"
)

var (
	_ PreconditionFailed = (*PreconditionFailedError)(nil)
	_ Error              = (*PreconditionFailedError)(nil)
)

type PreconditionFailed interface {
	error
	IsPreconditionFailed()
}

type PreconditionFailedError struct {
	*Zerror
}

func ThrowPreconditionFailed(parent error, id, message string) error {
	return &PreconditionFailedError{CreateZerror(parent, id, message)}
}

func ThrowPreconditionFailedf(parent error, id, format string, a ...interface{}) error {
	return ThrowPreconditionFailed(parent, id, fmt.Sprintf(format, a...))
}

func (err *PreconditionFailedError) IsPreconditionFailed() {}

func IsPreconditionFailed(err error) bool {
	_, ok := err.(PreconditionFailed)
	return ok
}

func (err *PreconditionFailedError) Is(target error) bool {
	t, ok := target.(*PreconditionFailedError)
	if !ok {
		return false
	}
	return err.Zerror.Is(t.Zerror)
}

func (err *PreconditionFailedError) Unwrap() error {
	return err.Zerror
}
