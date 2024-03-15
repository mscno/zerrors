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
