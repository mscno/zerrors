package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ PermissionDenied = (*PermissionDeniedError)(nil)
	_ Error            = (*PermissionDeniedError)(nil)
)

const PermissionDeniedId = "PermissionDenied"

type PermissionDenied interface {
	error
	IsPermissionDenied()
}

type PermissionDeniedError struct {
	*Zerror
}

func ThrowPermissionDenied(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return &PermissionDeniedError{CreateZerror(nil, PermissionDeniedId, message)}
}

func ThrowPermissionDeniedr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &PermissionDeniedError{CreateZerror(nil, PermissionDeniedId, message)}
}

func ToPermissionDenied(parent error, id string, message string) error {
	return &PermissionDeniedError{CreateZerror(parent, id, message)}
}

func ToPermissionDeniedf(parent error, id string, format string, a ...interface{}) error {
	return ToPermissionDenied(parent, id, fmt.Sprintf(format, a...))
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
