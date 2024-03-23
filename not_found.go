package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrNotFound = "not found"

func ThrowNotFound(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrNotFound).Errorf(message)
}

func ThrowNotFoundr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrNotFound).Errorf(message)
}

func NotFound(format string, a ...interface{}) error {
	return oops.Code(ErrNotFound).Errorf(format, a...)
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
