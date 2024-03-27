package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrInternal = "internal"

func ThrowInternal(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrInternal).Errorf(message)
}

func ThrowInternalr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrInternal).Errorf(message)
}

func Internal(message string, KVs ...interface{}) error {
	return oops.Code(ErrInternal).With(KVs...).Errorf(message)
}

func Internalf(format string, a ...interface{}) error {
	return oops.Code(ErrInternal).Errorf(format, a...)
}

func InternalBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrInternal)
}

func ToInternal(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrInternal).Wrapf(parent, format, a...)
}

func IsInternal(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrInternal
	}
	return false
}
