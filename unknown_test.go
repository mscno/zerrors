package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownError(t *testing.T) {
	var err interface{} = new(UnknownError)
	_, ok := err.(Unknown)
	assert.True(t, ok)
}

func TestThrowUnknownf(t *testing.T) {
	err := ThrowUnknownf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*UnknownError)
	assert.True(t, ok)
}

func TestIsUnknown(t *testing.T) {
	err := ThrowUnknown(nil, "id", "msg")
	ok := IsUnknown(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnknown(err)
	assert.False(t, ok)
}
