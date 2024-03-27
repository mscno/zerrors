package zerrors

import (
	"errors"
	"fmt"
	"github.com/samber/oops"
)

const ErrAlreadyExists = "already exists"

func ThrowAlreadyExists(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return oops.Code(ErrAlreadyExists).Errorf(message)
}

func ThrowAlreadyExistsr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return oops.Code(ErrAlreadyExists).Errorf(message)
}

func AlreadyExists(message string, KVs ...interface{}) error {
	return oops.Code(ErrAlreadyExists).With(KVs...).Errorf(message)
}

func AlreadyExistsf(format string, a ...interface{}) error {
	return oops.Code(ErrAlreadyExists).Errorf(format, a...)
}

func AlreadyExistsBuilder() oops.OopsErrorBuilder {
	return oops.Code(ErrAlreadyExists)
}

func ToAlreadyExists(parent error, format string, a ...interface{}) error {
	return oops.Code(ErrAlreadyExists).Wrapf(parent, format, a...)
}

func IsAlreadyExists(err error) bool {
	var possibleError oops.OopsError
	if errors.As(err, &possibleError) {
		return possibleError.Code() == ErrAlreadyExists
	}
	return false
}
