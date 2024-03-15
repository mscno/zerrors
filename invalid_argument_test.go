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
