package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalError(t *testing.T) {
	var err interface{} = new(InternalError)
	_, ok := err.(Internal)
	assert.True(t, ok)
}

func TestThrowInternalf(t *testing.T) {
	err := ThrowInternal(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*InternalError)
	assert.True(t, ok)
}

func TestIsInternal(t *testing.T) {
	err := ThrowInternal(nil, "id", "msg")
	ok := IsInternal(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsInternal(err)
	assert.False(t, ok)
}
