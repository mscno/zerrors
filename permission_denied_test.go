package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionDeniedError(t *testing.T) {
	var err interface{}
	err = new(PermissionDeniedError)
	_, ok := err.(PermissionDenied)
	assert.True(t, ok)
}

func TestThrowPermissionDeniedf(t *testing.T) {
	err := ThrowPermissionDeniedf(nil, "id", "msg")
	_, ok := err.(*PermissionDeniedError)
	assert.True(t, ok)
}

func TestIsPermissionDenied(t *testing.T) {
	err := ThrowPermissionDenied(nil, "id", "msg")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	ok = errors.Is(err, &PermissionDeniedError{})
	assert.True(t, ok)

	var e *PermissionDeniedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsPermissionDenied(err)
	assert.False(t, ok)

	ok = errors.Is(err, &PermissionDeniedError{})
	assert.False(t, ok)
}

func TestFindWrappedPermissionDenied(t *testing.T) {
	err := ThrowPermissionDenied(nil, "id", "msg")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsPermissionDenied(err)
	assert.True(t, ok)

	ok = errors.Is(err, &PermissionDeniedError{})
	assert.True(t, ok)

	var e *PermissionDeniedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}
