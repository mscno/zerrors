package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailedPreconditionError(t *testing.T) {
	var err interface{}
	err = new(FailedPreconditionError)
	_, ok := err.(FailedPrecondition)
	assert.True(t, ok)
}

func TestThrowFailedPreconditionf(t *testing.T) {
	err := ThrowFailedPreconditionf(nil, "id", "msg")
	_, ok := err.(*FailedPreconditionError)
	assert.True(t, ok)
}

func TestIsFailedPrecondition(t *testing.T) {
	err := ThrowFailedPrecondition(nil, "id", "msg")
	ok := IsFailedPrecondition(err)
	assert.True(t, ok)

	ok = errors.Is(err, &FailedPreconditionError{})
	assert.True(t, ok)

	var e *FailedPreconditionError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsFailedPrecondition(err)
	assert.False(t, ok)

	ok = errors.Is(err, &FailedPreconditionError{})
	assert.False(t, ok)
}

func TestFindWrappedFailedPrecondition(t *testing.T) {
	err := ThrowFailedPrecondition(nil, "id", "msg")
	ok := IsFailedPrecondition(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsFailedPrecondition(err)
	assert.True(t, ok)

	ok = errors.Is(err, &FailedPreconditionError{})
	assert.True(t, ok)

	var e *FailedPreconditionError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}
