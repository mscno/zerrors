package zerrors

import (
	"errors"
	"fmt"
	"reflect"
)

var _ Error = (*Zerror)(nil)

type Zerror struct {
	Parent  error
	Message string
	ID      string
}

func ThrowError(parent error, id, message string) error {
	return CreateZerror(parent, id, message)
}

func CreateZerror(parent error, id, message string) *Zerror {
	return &Zerror{
		Parent:  parent,
		ID:      id,
		Message: message,
	}
}

func (err *Zerror) Error() string {
	if err.Parent != nil {
		return fmt.Sprintf("ID=%s Message=%s Parent=(%v)", err.ID, err.Message, err.Parent)
	}
	return fmt.Sprintf("ID=%s Message=%s", err.ID, err.Message)
}

func (err *Zerror) Unwrap() error {
	return err.GetParent()
}

func (err *Zerror) GetParent() error {
	return err.Parent
}

func (err *Zerror) GetMessage() string {
	return err.Message
}

func (err *Zerror) SetMessage(msg string) {
	err.Message = msg
}

func (err *Zerror) GetID() string {
	return err.ID
}

func (err *Zerror) Is(target error) bool {
	t, ok := target.(*Zerror)
	if !ok {
		return false
	}
	if t.ID != "" && t.ID != err.ID {
		return false
	}
	if t.Message != "" && t.Message != err.Message {
		return false
	}
	if t.Parent != nil && !errors.Is(err.Parent, t.Parent) {
		return false
	}

	return true
}

func (err *Zerror) As(target interface{}) bool {
	_, ok := target.(**Zerror)
	if !ok {
		return false
	}
	reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
	return true
}
