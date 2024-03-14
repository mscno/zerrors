package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidArgumentError(t *testing.T) {
	var invalidArgumentError interface{} = new(InvalidArgumentError)
	_, ok := invalidArgumentError.(InvalidArgument)
	assert.True(t, ok)
}

func TestThrowInvalidArgumentf(t *testing.T) {
	err := ThrowInvalidArgumentf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*InvalidArgumentError)
	assert.True(t, ok)
}

func TestIsErrorInvalidArgument(t *testing.T) {
	err := ThrowInvalidArgument(nil, "id", "msg")
	ok := IsErrorInvalidArgument(err)
	assert.True(t, ok)

	err = errors.New("I am invalid!")
	ok = IsErrorInvalidArgument(err)
	assert.False(t, ok)
}
