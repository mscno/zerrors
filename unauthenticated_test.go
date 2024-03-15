package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnauthenticatedError(t *testing.T) {
	var err interface{}
	err = new(UnauthenticatedError)
	_, ok := err.(Unauthenticated)
	assert.True(t, ok)
}

func TestThrowUnauthenticatedf(t *testing.T) {
	err := ThrowUnauthenticatedf(nil, "id", "msg")
	_, ok := err.(*UnauthenticatedError)
	assert.True(t, ok)
}

func TestIsUnauthenticated(t *testing.T) {
	err := ThrowUnauthenticated(nil, "id", "msg")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnauthenticatedError{})
	assert.True(t, ok)

	var e *UnauthenticatedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnauthenticated(err)
	assert.False(t, ok)

	ok = errors.Is(err, &UnauthenticatedError{})
	assert.False(t, ok)
}

func TestFindWrappedUnauthenticated(t *testing.T) {
	err := ThrowUnauthenticated(nil, "id", "msg")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsUnauthenticated(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnauthenticatedError{})
	assert.True(t, ok)

	var e *UnauthenticatedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}
