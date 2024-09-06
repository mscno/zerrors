package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrUnavailable = "unavailable"

func ThrowUnavailable(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrUnavailable).Wrap(errors.New(message))
}

func ThrowUnavailabler(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrUnavailable).Wrap(errors.New(message))
}

func Unavailable(message string, KVs ...interface{}) error {
	return oops.Code(ErrUnavailable).With(KVs...).Wrap(errors.New(message))
}

func Unavailablef(format string, a ...interface{}) error {
	return oops.Code(ErrUnavailable).Wrap(fmt.Errorf(format, a...))
}

func UnavailableBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrUnavailable)
}

func ToUnavailable(parent error, KVs ...interface{}) error {
	return oops.Code(ErrUnavailable).With(KVs...).Wrap(parent)
}

func ToUnavailablef(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrUnavailable).Wrapf(parent, format, a...)
}

func IsUnavailable(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrUnavailable
	}
	return false
}
