package zerrors

import "fmt"

var (
	_ InvalidArgument = (*InvalidArgumentError)(nil)
	_ Error           = (*InvalidArgumentError)(nil)
)

type InvalidArgument interface {
	error
	IsInvalidArgument()
}

type InvalidArgumentError struct {
	*Zerror
}

func ThrowInvalidArgument(parent error, id, message string) error {
	return &InvalidArgumentError{CreateZnowError(parent, id, message)}
}

func ThrowInvalidArgumentf(parent error, id, format string, a ...interface{}) error {
	return ThrowInvalidArgument(parent, id, fmt.Sprintf(format, a...))
}

func (err *InvalidArgumentError) IsInvalidArgument() {}

func IsErrorInvalidArgument(err error) bool {
	_, ok := err.(InvalidArgument)
	return ok
}

func (err *InvalidArgumentError) Is(target error) bool {
	t, ok := target.(*InvalidArgumentError)
	if !ok {
		return false
	}
	return err.Zerror.Is(t.Zerror)
}

func (err *InvalidArgumentError) Unwrap() error {
	return err.Zerror
}
