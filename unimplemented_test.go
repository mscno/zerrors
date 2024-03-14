package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnimplementedError(t *testing.T) {
	var unimplementedError interface{} = new(UnimplementedError)
	_, ok := unimplementedError.(Unimplemented)
	assert.True(t, ok)
}

func TestThrowUnimplementedf(t *testing.T) {
	err := ThrowUnimplementedf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*UnimplementedError)
	assert.True(t, ok)
}

func TestIsUnimplemented(t *testing.T) {
	err := ThrowUnimplemented(nil, "id", "msg")
	ok := IsUnimplemented(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnimplemented(err)
	assert.False(t, ok)
}
