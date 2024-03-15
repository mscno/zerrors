package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeadlineExceededError(t *testing.T) {
	var err interface{}
	err = new(DeadlineExceededError)
	_, ok := err.(DeadlineExceeded)
	assert.True(t, ok)
}

func TestThrowDeadlineExceededf(t *testing.T) {
	err := ThrowDeadlineExceededf(nil, "id", "msg")
	_, ok := err.(*DeadlineExceededError)
	assert.True(t, ok)
}

func TestIsDeadlineExceeded(t *testing.T) {
	err := ThrowDeadlineExceeded(nil, "id", "msg")
	ok := IsDeadlineExceeded(err)
	assert.True(t, ok)

	ok = errors.Is(err, &DeadlineExceededError{})
	assert.True(t, ok)

	var e *DeadlineExceededError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsDeadlineExceeded(err)
	assert.False(t, ok)

	ok = errors.Is(err, &DeadlineExceededError{})
	assert.False(t, ok)
}

func TestFindWrappedDeadlineExceeded(t *testing.T) {
	err := ThrowDeadlineExceeded(nil, "id", "msg")
	ok := IsDeadlineExceeded(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsDeadlineExceeded(err)
	assert.True(t, ok)

	ok = errors.Is(err, &DeadlineExceededError{})
	assert.True(t, ok)

	var e *DeadlineExceededError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}
