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

func TestToFailedPreconditionf(t *testing.T) {
	err := ToFailedPreconditionf(nil, "id", "msg")
	_, ok := err.(*FailedPreconditionError)
	assert.True(t, ok)
}

func TestIsFailedPrecondition(t *testing.T) {
	err := ToFailedPrecondition(nil, "id", "msg")
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
	err := ToFailedPrecondition(nil, "id", "msg")
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

func TestFailedPreconditionWithRootCause(t *testing.T) {
	err := ToFailedPrecondition(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsFailedPrecondition(err)
	assert.True(t, ok)

	ok = errors.Is(err, &FailedPreconditionError{})
	assert.True(t, ok)

	var asDomain *FailedPreconditionError
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

func TestWrappingAnotherFailedPreconditionError(t *testing.T) {
	err := ToFailedPrecondition(nil, "id1", "msg")
	ok := IsFailedPrecondition(err)
	assert.True(t, ok)

	err = ToFailedPrecondition(err, "id2", "msg")
	ok = IsFailedPrecondition(err)
	assert.True(t, ok)

	ok = errors.Is(err, &FailedPreconditionError{})
	assert.True(t, ok)

	var e *FailedPreconditionError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
