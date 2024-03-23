package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrUnauthenticated = "unauthenticated"

func ThrowUnauthenticated(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrUnauthenticated).Errorf(message)
}

func ThrowUnauthenticatedr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrUnauthenticated).Errorf(message)
}

func Unauthenticated(format string, a ...interface{}) error {
	return oops.Code(ErrUnauthenticated).Errorf(format, a...)
}

func ToUnauthenticated(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrUnauthenticated).Wrapf(parent, format, a...)
}

func IsUnauthenticated(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrUnauthenticated
	}
	return false
}
