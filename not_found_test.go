package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundError(t *testing.T) {
	var notFoundError interface{} = new(NotFoundError)
	_, ok := notFoundError.(NotFound)
	assert.True(t, ok)
}

func TestThrowNotFoundf(t *testing.T) {
	err := ThrowNotFoundf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*NotFoundError)
	assert.True(t, ok)
}

func TestIsNotFound(t *testing.T) {
	err := ThrowNotFound(nil, "id", "msg")
	ok := IsNotFound(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsNotFound(err)
	assert.False(t, ok)
}
