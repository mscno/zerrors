package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnauthenticatedError(t *testing.T) {
	var err interface{} = new(UnauthenticatedError)
	_, ok := err.(Unauthenticated)
	assert.True(t, ok)
}

func TestThrowUnauthenticatedf(t *testing.T) {
	err := ThrowUnauthenticatedf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*UnauthenticatedError)
	assert.True(t, ok)
}

func TestIsUnauthenticated(t *testing.T) {
	err := ThrowUnauthenticated(nil, "id", "msg")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnauthenticated(err)
	assert.False(t, ok)
}
