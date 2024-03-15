package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ PermissionDenied = (*PermissionDeniedError)(nil)
	_ Error            = (*PermissionDeniedError)(nil)
)

type PermissionDenied interface {
	error
	IsPermissionDenied()
}

type PermissionDeniedError struct {
	*Zerror
}

func ThrowPermissionDenied(parent error, id string, message string) error {
	return &PermissionDeniedError{CreateZerror(parent, id, message)}
}

func ThrowPermissionDeniedf(parent error, id string, format string, a ...interface{}) error {
	return ThrowPermissionDenied(parent, id, fmt.Sprintf(format, a...))
}

func (err *PermissionDeniedError) IsPermissionDenied() {}

func (err *PermissionDeniedError) Is(target error) bool {
	return IsPermissionDenied(target)
}

func IsPermissionDenied(err error) bool {
	var possibleError *PermissionDeniedError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
