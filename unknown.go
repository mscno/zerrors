package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrUnknown = "unknown"

func ThrowUnknown(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrUnknown).Errorf(message)
}

func ThrowUnknownr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrUnknown).Errorf(message)
}

func Unknown(format string, a ...interface{}) error {
	return oops.Code(ErrUnknown).Errorf(format, a...)
}

func UnknownBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrUnknown)
}

func ToUnknown(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrUnknown).Wrapf(parent, format, a...)
}

func IsUnknown(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrUnknown
	}
	return false
}
