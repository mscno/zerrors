package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrFailedPrecondition = "failed precondition"

func ThrowFailedPrecondition(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrFailedPrecondition).Errorf(message)
}

func ThrowFailedPreconditionr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrFailedPrecondition).Errorf(message)
}

func FailedPrecondition(message string, KVs ...interface{}) error {
	return oops.Code(ErrFailedPrecondition).With(KVs...).Errorf(message)
}

func FailedPreconditionf(format string, a ...interface{}) error {
	return oops.Code(ErrFailedPrecondition).Errorf(format, a...)
}

func FailedPreconditionBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrFailedPrecondition)
}

func ToFailedPrecondition(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrFailedPrecondition).Wrapf(parent, format, a...)
}

func IsFailedPrecondition(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrFailedPrecondition
	}
	return false
}
