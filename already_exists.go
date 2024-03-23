package zerrors

import (
	"errors"
	"fmt"
)

var (
	_ AlreadyExists = (*AlreadyExistsError)(nil)
	_ Error         = (*AlreadyExistsError)(nil)
)

const AlreadyExistsId = "AlreadyExists"

type AlreadyExists interface {
	error
	IsAlreadyExists()
}

type AlreadyExistsError struct {
	*Zerror
}

func ThrowAlreadyExists(action, kind, name, reason string) error {
	message := fmt.Sprintf("cannot %s '%s' of kind '%s': %s", action, name, kind, reason)
	return &AlreadyExistsError{CreateZerror(nil, AlreadyExistsId, message)}
}

func ToAlreadyExists(parent error, id string, message string) error {
	return &AlreadyExistsError{CreateZerror(parent, id, message)}
}

func ToAlreadyExistsf(parent error, id string, format string, a ...interface{}) error {
	return ToAlreadyExists(parent, id, fmt.Sprintf(format, a...))
}

func (err *AlreadyExistsError) IsAlreadyExists() {}

func (err *AlreadyExistsError) Is(target error) bool {
	return IsAlreadyExists(target)
}

func IsAlreadyExists(err error) bool {
	var possibleError *AlreadyExistsError
	if errors.As(err, &possibleError) {
		return true
	}
	return false
}
