package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrUnauthenticated = "unauthenticated"

func ThrowUnauthenticated(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrUnauthenticated).Wrap(errors.New(message))
}

func ThrowUnauthenticatedr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrUnauthenticated).Wrap(errors.New(message))
}

func Unauthenticated(message string, KVs ...interface{}) error {
	return oops.Code(ErrUnauthenticated).With(KVs...).Wrap(errors.New(message))
}

func Unauthenticatedf(format string, a ...interface{}) error {
	return oops.Code(ErrUnauthenticated).Wrap(fmt.Errorf(format, a...))
}

func UnauthenticatedBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrUnauthenticated)
}

func ToUnauthenticated(parent error, KVs ...interface{}) error {
	return oops.Code(ErrUnauthenticated).With(KVs...).Wrap(parent)
}

func ToUnauthenticatedf(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrUnauthenticated).Wrapf(parent, format, a...)
}

func IsUnauthenticated(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrUnauthenticated
	}
	return false
}
