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

func TestThrowNotFoundf(t *testing.T) {
	err := ThrowNotFoundf(nil, "id", "msg")
	_, ok := err.(*NotFoundError)
	assert.True(t, ok)
}

func TestIsNotFound(t *testing.T) {
	err := ThrowNotFound(nil, "id", "msg")
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
	err := ThrowNotFound(nil, "id", "msg")
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
