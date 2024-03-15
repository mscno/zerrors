package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnimplementedError(t *testing.T) {
	var err interface{}
	err = new(UnimplementedError)
	_, ok := err.(Unimplemented)
	assert.True(t, ok)
}

func TestThrowUnimplementedf(t *testing.T) {
	err := ThrowUnimplementedf(nil, "id", "msg")
	_, ok := err.(*UnimplementedError)
	assert.True(t, ok)
}

func TestIsUnimplemented(t *testing.T) {
	err := ThrowUnimplemented(nil, "id", "msg")
	ok := IsUnimplemented(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnimplementedError{})
	assert.True(t, ok)

	var e *UnimplementedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnimplemented(err)
	assert.False(t, ok)

	ok = errors.Is(err, &UnimplementedError{})
	assert.False(t, ok)
}

func TestFindWrappedUnimplemented(t *testing.T) {
	err := ThrowUnimplemented(nil, "id", "msg")
	ok := IsUnimplemented(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsUnimplemented(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnimplementedError{})
	assert.True(t, ok)

	var e *UnimplementedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}
