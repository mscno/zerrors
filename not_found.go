package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrNotFound = "not found"

func ThrowNotFound(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrNotFound).Wrap(errors.New(message))
}

func ThrowNotFoundr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrNotFound).Wrap(errors.New(message))
}

func NotFound(message string, KVs ...interface{}) error {
	return oops.Code(ErrNotFound).With(KVs...).Wrap(errors.New(message))
}

func NotFoundf(format string, a ...interface{}) error {
	return oops.Code(ErrNotFound).Wrap(fmt.Errorf(format, a...))
}

func NotFoundBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrNotFound)
}

func ToNotFound(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrNotFound).Wrapf(parent, format, a...)
}

func IsNotFound(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrNotFound
	}
	return false
}
