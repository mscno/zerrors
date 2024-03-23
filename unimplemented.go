package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Unimplemented = (*UnimplementedError)(nil)
	_ Error         = (*UnimplementedError)(nil)
)

const UnimplementedId = "Unimplemented"

type Unimplemented interface {
	error
	IsUnimplemented()
}

type UnimplementedError struct {
	*Zerror
}

func ThrowUnimplemented(action, kind, name string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s'", action, name, kind)
	return &UnimplementedError{CreateZerror(nil, UnimplementedId, message)}
}

func ThrowUnimplementedr(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &UnimplementedError{CreateZerror(nil, UnimplementedId, message)}
}

func ToUnimplemented(parent error, id string, message string) error {
	return &UnimplementedError{CreateZerror(parent, id, message)}
}

func ToUnimplementedf(parent error, id string, format string, a ...interface{}) error {
	return ToUnimplemented(parent, id, fmt.Sprintf(format, a...))
}

func (err *UnimplementedError) IsUnimplemented() {}

func (err *UnimplementedError) Is(target error) bool {
	return IsUnimplemented(target)
}

func IsUnimplemented(err error) bool {
	var possibleError *UnimplementedError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
