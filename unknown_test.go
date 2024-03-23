package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownError(t *testing.T) {
	var err interface{}
	err = new(UnknownError)
	_, ok := err.(Unknown)
	assert.True(t, ok)
}

func TestThrowUnknownf(t *testing.T) {
	err := ThrowUnknownf(nil, "id", "msg")
	_, ok := err.(*UnknownError)
	assert.True(t, ok)
}

func TestIsUnknown(t *testing.T) {
	err := ThrowUnknown(nil, "id", "msg")
	ok := IsUnknown(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnknownError{})
	assert.True(t, ok)

	var e *UnknownError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnknown(err)
	assert.False(t, ok)

	ok = errors.Is(err, &UnknownError{})
	assert.False(t, ok)
}

func TestFindWrappedUnknown(t *testing.T) {
	err := ThrowUnknown(nil, "id", "msg")
	ok := IsUnknown(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsUnknown(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnknownError{})
	assert.True(t, ok)

	var e *UnknownError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestUnknownWithRootCause(t *testing.T) {
	err := ThrowUnknown(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsUnknown(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnknownError{})
	assert.True(t, ok)

	var asDomain *UnknownError
	ok = errors.As(err, &asDomain)
	assert.True(t, ok)

	if asDomain.Zerror.Parent == nil {
		t.Fatal("underlying cause was not preserved")
	}

	assert.Equal(t, "not domain", asDomain.Zerror.Parent.Error())

	assert.Equalf(t, "id", asDomain.Zerror.ID, "ID is not equal")
	assert.Equalf(t, "no wrap message", asDomain.Zerror.Message, "Message is not equal")

	msg := "ID=id Message=no wrap message Parent=(not domain)"
	assert.Equal(t, msg, err.Error())

}

func TestWrappingAnotherUnknownError(t *testing.T) {
	err := ThrowUnknown(nil, "id1", "msg")
	ok := IsUnknown(err)
	assert.True(t, ok)

	err = ThrowUnknown(err, "id2", "msg")
	ok = IsUnknown(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnknownError{})
	assert.True(t, ok)

	var e *UnknownError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
