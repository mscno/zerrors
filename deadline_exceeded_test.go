package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeadlineExceededError(t *testing.T) {
	var err interface{} = new(DeadlineExceededError)
	_, ok := err.(DeadlineExceeded)
	assert.True(t, ok)
}

func TestThrowDeadlineExceededf(t *testing.T) {
	err := ThrowDeadlineExceededf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*DeadlineExceededError)
	assert.True(t, ok)
}

func TestIsDeadlineExceeded(t *testing.T) {
	err := ThrowDeadlineExceeded(nil, "id", "msg")
	ok := IsDeadlineExceeded(err)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsDeadlineExceeded(err)
	assert.False(t, ok)
}
