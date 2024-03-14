package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnavailableError(t *testing.T) {
	var err interface{} = new(UnavailableError)
	_, ok := err.(Unavailable)
	assert.True(t, ok)
}

func TestThrowUnavailablef(t *testing.T) {
	err := ThrowUnavailablef(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*UnavailableError)
	assert.True(t, ok)
}

func TestIsUnavailable(t *testing.T) {
	err := ThrowUnavailable(nil, "id", "msg")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnavailable(err)
	assert.False(t, ok)
}
