package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidArgumentError(t *testing.T) {
	var err interface{}
	err = new(InvalidArgumentError)
	_, ok := err.(InvalidArgument)
	assert.True(t, ok)
}

func TestThrowInvalidArgumentf(t *testing.T) {
	err := ThrowInvalidArgumentf(nil, "id", "msg")
	_, ok := err.(*InvalidArgumentError)
	assert.True(t, ok)
}

func TestIsInvalidArgument(t *testing.T) {
	err := ThrowInvalidArgument(nil, "id", "msg")
	ok := IsInvalidArgument(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InvalidArgumentError{})
	assert.True(t, ok)

	var e *InvalidArgumentError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsInvalidArgument(err)
	assert.False(t, ok)

	ok = errors.Is(err, &InvalidArgumentError{})
	assert.False(t, ok)
}

func TestFindWrappedInvalidArgument(t *testing.T) {
	err := ThrowInvalidArgument(nil, "id", "msg")
	ok := IsInvalidArgument(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsInvalidArgument(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InvalidArgumentError{})
	assert.True(t, ok)

	var e *InvalidArgumentError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestInvalidArgumentWithRootCause(t *testing.T) {
	err := ThrowInvalidArgument(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsInvalidArgument(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InvalidArgumentError{})
	assert.True(t, ok)

	var asDomain *InvalidArgumentError
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

func TestWrappingAnotherInvalidArgumentError(t *testing.T) {
	err := ThrowInvalidArgument(nil, "id1", "msg")
	ok := IsInvalidArgument(err)
	assert.True(t, ok)

	err = ThrowInvalidArgument(err, "id2", "msg")
	ok = IsInvalidArgument(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InvalidArgumentError{})
	assert.True(t, ok)

	var e *InvalidArgumentError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
