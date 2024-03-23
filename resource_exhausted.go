package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrResourceExhausted = "resource exhausted"

func ThrowResourceExhausted(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrResourceExhausted).Errorf(message)
}

func ThrowResourceExhaustedr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrResourceExhausted).Errorf(message)
}

func ResourceExhausted(format string, a ...interface{}) error {
	return oops.Code(ErrResourceExhausted).Errorf(format, a...)
}

func ToResourceExhausted(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrResourceExhausted).Wrapf(parent, format, a...)
}

func IsResourceExhausted(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrResourceExhausted
	}
	return false
}
