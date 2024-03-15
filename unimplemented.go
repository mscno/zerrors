package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ Unimplemented = (*UnimplementedError)(nil)
	_ Error         = (*UnimplementedError)(nil)
)

type Unimplemented interface {
	error
	IsUnimplemented()
}

type UnimplementedError struct {
	*Zerror
}

func ThrowUnimplemented(parent error, id string, message string) error {
	return &UnimplementedError{CreateZerror(parent, id, message)}
}

func ThrowUnimplementedf(parent error, id string, format string, a ...interface{}) error {
	return ThrowUnimplemented(parent, id, fmt.Sprintf(format, a...))
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
