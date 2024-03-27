package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrUnimplemented = "unimplemented"

func ThrowUnimplemented(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrUnimplemented).Errorf(message)
}

func ThrowUnimplementedr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrUnimplemented).Errorf(message)
}

func Unimplemented(message string, KVs ...interface{}) error {
	return oops.Code(ErrUnimplemented).With(KVs...).Errorf(message)
}

func Unimplementedf(format string, a ...interface{}) error {
	return oops.Code(ErrUnimplemented).Errorf(format, a...)
}

func UnimplementedBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrUnimplemented)
}

func ToUnimplemented(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrUnimplemented).Wrapf(parent, format, a...)
}

func IsUnimplemented(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrUnimplemented
	}
	return false
}
