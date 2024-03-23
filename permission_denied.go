package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrPermissionDenied = "permission denied"

func ThrowPermissionDenied(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrPermissionDenied).Errorf(message)
}

func ThrowPermissionDeniedr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrPermissionDenied).Errorf(message)
}

func PermissionDenied(format string, a ...interface{}) error {
	return oops.Code(ErrPermissionDenied).Errorf(format, a...)
}

func PermissionDeniedBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrPermissionDenied)
}

func ToPermissionDenied(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrPermissionDenied).Wrapf(parent, format, a...)
}

func IsPermissionDenied(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrPermissionDenied
	}
	return false
}
