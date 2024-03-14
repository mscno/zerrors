package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceExhaustedError(t *testing.T) {
	var err interface{} = new(ResourceExhaustedError)
	_, ok := err.(ResourceExhausted)
	assert.True(t, ok)
}

func TestThrowResourceExhaustedf(t *testing.T) {
	err := ThrowResourceExhaustedf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*ResourceExhaustedError)
	assert.True(t, ok)
}

func TestIsResourceExhausted(t *testing.T) {
	err := ThrowResourceExhausted(nil, "id", "msg")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsResourceExhausted(err)
	assert.False(t, ok)
}
