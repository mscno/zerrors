package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceExhaustedError(t *testing.T) {
	var err interface{}
	err = new(ResourceExhaustedError)
	_, ok := err.(ResourceExhausted)
	assert.True(t, ok)
}

func TestThrowResourceExhaustedf(t *testing.T) {
	err := ThrowResourceExhaustedf(nil, "id", "msg")
	_, ok := err.(*ResourceExhaustedError)
	assert.True(t, ok)
}

func TestIsResourceExhausted(t *testing.T) {
	err := ThrowResourceExhausted(nil, "id", "msg")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	ok = errors.Is(err, &ResourceExhaustedError{})
	assert.True(t, ok)

	var e *ResourceExhaustedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsResourceExhausted(err)
	assert.False(t, ok)

	ok = errors.Is(err, &ResourceExhaustedError{})
	assert.False(t, ok)
}

func TestFindWrappedResourceExhausted(t *testing.T) {
	err := ThrowResourceExhausted(nil, "id", "msg")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsResourceExhausted(err)
	assert.True(t, ok)

	ok = errors.Is(err, &ResourceExhaustedError{})
	assert.True(t, ok)

	var e *ResourceExhaustedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}
