package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionDeniedError(t *testing.T) {
	var err interface{} = new(PermissionDeniedError)
	_, ok := err.(PermissionDenied)
	assert.True(t, ok)
}

func TestThrowPermissionDeniedf(t *testing.T) {
	err := ThrowPermissionDeniedf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*PermissionDeniedError)
	assert.True(t, ok)
}

func TestIsPermissionDenied(t *testing.T) {
	err := ThrowPermissionDenied(nil, "id", "msg")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsPermissionDenied(err)
	assert.False(t, ok)
}
