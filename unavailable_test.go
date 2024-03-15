package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnavailableError(t *testing.T) {
	var err interface{}
	err = new(UnavailableError)
	_, ok := err.(Unavailable)
	assert.True(t, ok)
}

func TestThrowUnavailablef(t *testing.T) {
	err := ThrowUnavailablef(nil, "id", "msg")
	_, ok := err.(*UnavailableError)
	assert.True(t, ok)
}

func TestIsUnavailable(t *testing.T) {
	err := ThrowUnavailable(nil, "id", "msg")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnavailableError{})
	assert.True(t, ok)

	var e *UnavailableError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnavailable(err)
	assert.False(t, ok)

	ok = errors.Is(err, &UnavailableError{})
	assert.False(t, ok)
}

func TestFindWrappedUnavailable(t *testing.T) {
	err := ThrowUnavailable(nil, "id", "msg")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsUnavailable(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnavailableError{})
	assert.True(t, ok)

	var e *UnavailableError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}
