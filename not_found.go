package zerrors

import "fmt"

type NotFound interface {
	error
	IsNotFound()
}

type NotFoundError struct {
	*Zerror
}

func ThrowNotFound(parent error, id, message string) error {
	return &NotFoundError{CreateZerror(parent, id, message)}
}

func ThrowNotFoundf(parent error, id, format string, a ...interface{}) error {
	return ThrowNotFound(parent, id, fmt.Sprintf(format, a...))
}

func ThrowNotFoundResource(action, kind, name string) error {
	return ThrowNotFoundf(nil, action, "cannot %s %s [%s]", action, kind, name)
}

func ThrowNotFoundResourceParent(action, kind, parent, parentKind string) error {
	return ThrowNotFoundf(nil, action, "cannot %s %s, parent %s [%s] does not exist", action, kind, parentKind, parent)
}

func (err *NotFoundError) IsNotFound() {}

func IsNotFound(err error) bool {
	_, ok := err.(NotFound)
	return ok
}

func (err *NotFoundError) Is(target error) bool {
	t, ok := target.(*NotFoundError)
	if !ok {
		return false
	}
	return err.Zerror.Is(t.Zerror)
}

func (err *NotFoundError) Unwrap() error {
	return err.Zerror
}
