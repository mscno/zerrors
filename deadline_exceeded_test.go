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

func TestDeadlineExceededWithRootCause(t *testing.T) {
	err := ThrowDeadlineExceeded(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsDeadlineExceeded(err)
	assert.True(t, ok)

	ok = errors.Is(err, &DeadlineExceededError{})
	assert.True(t, ok)

	var asDomain *DeadlineExceededError
	ok = errors.As(err, &asDomain)
	assert.True(t, ok)

	if asDomain.Zerror.Parent == nil {
		t.Fatal("underlying cause was not preserved")
	}

	assert.Equal(t, "not domain", asDomain.Zerror.Parent.Error())

	assert.Equalf(t, "id", asDomain.Zerror.ID, "ID is not equal")
	assert.Equalf(t, "no wrap message", asDomain.Zerror.Message, "Message is not equal")

	msg := "ID=id Message=no wrap message Parent=(not domain)"
	assert.Equal(t, msg, err.Error())

}

func TestWrappingAnotherDeadlineExceededError(t *testing.T) {
	err := ThrowDeadlineExceeded(nil, "id1", "msg")
	ok := IsDeadlineExceeded(err)
	assert.True(t, ok)

	err = ThrowDeadlineExceeded(err, "id2", "msg")
	ok = IsDeadlineExceeded(err)
	assert.True(t, ok)

	ok = errors.Is(err, &DeadlineExceededError{})
	assert.True(t, ok)

	var e *DeadlineExceededError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
