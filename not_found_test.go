package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundError(t *testing.T) {
	var err interface{}
	err = new(NotFoundError)
	_, ok := err.(NotFound)
	assert.True(t, ok)
}

func TestToNotFoundf(t *testing.T) {
	err := ToNotFoundf(nil, "id", "msg")
	_, ok := err.(*NotFoundError)
	assert.True(t, ok)
}

func TestIsNotFound(t *testing.T) {
	err := ToNotFound(nil, "id", "msg")
	ok := IsNotFound(err)
	assert.True(t, ok)

	ok = errors.Is(err, &NotFoundError{})
	assert.True(t, ok)

	var e *NotFoundError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsNotFound(err)
	assert.False(t, ok)

	ok = errors.Is(err, &NotFoundError{})
	assert.False(t, ok)
}

func TestFindWrappedNotFound(t *testing.T) {
	err := ToNotFound(nil, "id", "msg")
	ok := IsNotFound(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsNotFound(err)
	assert.True(t, ok)

	ok = errors.Is(err, &NotFoundError{})
	assert.True(t, ok)

	var e *NotFoundError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestNotFoundWithRootCause(t *testing.T) {
	err := ToNotFound(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsNotFound(err)
	assert.True(t, ok)

	ok = errors.Is(err, &NotFoundError{})
	assert.True(t, ok)

	var asDomain *NotFoundError
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

func TestWrappingAnotherNotFoundError(t *testing.T) {
	err := ToNotFound(nil, "id1", "msg")
	ok := IsNotFound(err)
	assert.True(t, ok)

	err = ToNotFound(err, "id2", "msg")
	ok = IsNotFound(err)
	assert.True(t, ok)

	ok = errors.Is(err, &NotFoundError{})
	assert.True(t, ok)

	var e *NotFoundError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
