package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrDeadlineExceeded = "deadline exceeded"

func ThrowDeadlineExceeded(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrDeadlineExceeded).Wrap(errors.New(message))
}

func ThrowDeadlineExceededr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrDeadlineExceeded).Wrap(errors.New(message))
}

func DeadlineExceeded(message string, KVs ...interface{}) error {
	return oops.Code(ErrDeadlineExceeded).With(KVs...).Wrap(errors.New(message))
}

func DeadlineExceededf(format string, a ...interface{}) error {
	return oops.Code(ErrDeadlineExceeded).Wrap(fmt.Errorf(format, a...))
}

func DeadlineExceededBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrDeadlineExceeded)
}

func ToDeadlineExceeded(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrDeadlineExceeded).Wrapf(parent, format, a...)
}

func IsDeadlineExceeded(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrDeadlineExceeded
	}
	return false
}
