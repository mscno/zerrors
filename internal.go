package zerrors

import (
	"github.com/samber/oops"
)

var (
	_ Internal = (*InternalError)(nil)
	_ Error    = (*InternalError)(nil)
)

type Internal interface {
	error
	IsInternal()
}

type InternalError struct {
	*Zerror
}

func ThrowInternal(parent error, id, message string, kv ...any) error {
	o := oops.With(kv...)
	return &InternalError{CreateZerror(o.Wrap(parent), id, message)}
}

//func ThrowInternalf(parent error, id, format string, a ...interface{}) error {
//	return ThrowInternal(parent, id, fmt.Sprintf(format, a...))
//}

func (err *InternalError) IsInternal() {}

func IsInternal(err error) bool {
	_, ok := err.(Internal)
	return ok
}

func (err *InternalError) Is(target error) bool {
	t, ok := target.(*InternalError)
	if !ok {
		return false
	}
	return err.Zerror.Is(t.Zerror)
}

func (err *InternalError) Unwrap() error {
	return err.Zerror
}
