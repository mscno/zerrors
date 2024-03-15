package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlreadyExistsError(t *testing.T) {
	var err interface{}
	err = new(AlreadyExistsError)
	_, ok := err.(AlreadyExists)
	assert.True(t, ok)
}

func TestThrowAlreadyExistsf(t *testing.T) {
	err := ThrowAlreadyExistsf(nil, "id", "msg")
	_, ok := err.(*AlreadyExistsError)
	assert.True(t, ok)
}

func TestIsAlreadyExists(t *testing.T) {
	err := ThrowAlreadyExists(nil, "id", "msg")
	ok := IsAlreadyExists(err)
	assert.True(t, ok)

	ok = errors.Is(err, &AlreadyExistsError{})
	assert.True(t, ok)

	var e *AlreadyExistsError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	msg := "ID=id Message=msg"
	assert.Equal(t, msg, err.Error())

	err = errors.New("I am found!")
	ok = IsAlreadyExists(err)
	assert.False(t, ok)

	ok = errors.Is(err, &AlreadyExistsError{})
	assert.False(t, ok)
}

func TestFindWrappedAlreadyExists(t *testing.T) {
	err := ThrowAlreadyExists(nil, "id", "msg")
	ok := IsAlreadyExists(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsAlreadyExists(err)
	assert.True(t, ok)

	ok = errors.Is(err, &AlreadyExistsError{})
	assert.True(t, ok)

	var e *AlreadyExistsError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	msg := "wrapped ID=id Message=msg"
	assert.Equal(t, msg, err.Error())
}

func TestAlreadyExistsWithRootCause(t *testing.T) {
	err := ThrowAlreadyExists(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsAlreadyExists(err)
	assert.True(t, ok)

	ok = errors.Is(err, &AlreadyExistsError{})
	assert.True(t, ok)

	var asDomain *AlreadyExistsError
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

func TestWrappingAnotherAlreadyExistError(t *testing.T) {
	err := ThrowAlreadyExists(nil, "id", "msg")
	ok := IsAlreadyExists(err)
	assert.True(t, ok)

	err = ThrowAlreadyExists(err, "id", "msg")
	ok = IsAlreadyExists(err)
	assert.True(t, ok)

	ok = errors.Is(err, &AlreadyExistsError{})
	assert.True(t, ok)

	err = errors.New("I am found!")
	err = fmt.Errorf("wrapped: %w", err)
	err = fmt.Errorf("wrapped again: %w", err)

	var e *AlreadyExistsError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	msg := "ID=id Message=msg"
	assert.Equal(t, msg, err.Error())
}
