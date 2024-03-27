package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrInvalidArgument = "invalid argument"

func ThrowInvalidArgument(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrInvalidArgument).Errorf(message)
}

func ThrowInvalidArgumentr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrInvalidArgument).Errorf(message)
}

func InvalidArgument(message string, KVs ...interface{}) error {
	return oops.Code(ErrInvalidArgument).With(KVs...).Errorf(message)
}

func InvalidArgumentf(format string, a ...interface{}) error {
	return oops.Code(ErrInvalidArgument).Errorf(format, a...)
}

func InvalidArgumentBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrInvalidArgument)
}

func ToInvalidArgument(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrInvalidArgument).Wrapf(parent, format, a...)
}

func IsInvalidArgument(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrInvalidArgument
	}
	return false
}
