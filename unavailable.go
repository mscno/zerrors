package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrUnavailable = "unavailable"

func ThrowUnavailable(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrUnavailable).Errorf(message)
}

func ThrowUnavailabler(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrUnavailable).Errorf(message)
}

func Unavailable(message string, KVs ...interface{}) error {
	return oops.Code(ErrUnavailable).With(KVs...).Errorf(message)
}

func Unavailablef(format string, a ...interface{}) error {
	return oops.Code(ErrUnavailable).Errorf(format, a...)
}

func UnavailableBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrUnavailable)
}

func ToUnavailable(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrUnavailable).Wrapf(parent, format, a...)
}

func IsUnavailable(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrUnavailable
	}
	return false
}
