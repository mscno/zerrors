package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalError(t *testing.T) {
	var err interface{}
	err = new(InternalError)
	_, ok := err.(Internal)
	assert.True(t, ok)
}

func TestThrowInternalf(t *testing.T) {
	err := ThrowInternalf(nil, "id", "msg")
	_, ok := err.(*InternalError)
	assert.True(t, ok)
}

func TestIsInternal(t *testing.T) {
	err := ThrowInternal(nil, "id", "msg")
	ok := IsInternal(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InternalError{})
	assert.True(t, ok)

	var e *InternalError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsInternal(err)
	assert.False(t, ok)

	ok = errors.Is(err, &InternalError{})
	assert.False(t, ok)
}

func TestFindWrappedInternal(t *testing.T) {
	err := ThrowInternal(nil, "id", "msg")
	ok := IsInternal(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsInternal(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InternalError{})
	assert.True(t, ok)

	var e *InternalError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}
