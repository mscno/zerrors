package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalError(t *testing.T) {
	var err interface{}
	err = new(InternalError)
	_, ok := err.(Internal)
	assert.True(t, ok)
}

func TestThrowInternalf(t *testing.T) {
	err := ThrowInternalf(nil, "id", "msg")
	_, ok := err.(*InternalError)
	assert.True(t, ok)
}

func TestIsInternal(t *testing.T) {
	err := ThrowInternal(nil, "id", "msg")
	ok := IsInternal(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InternalError{})
	assert.True(t, ok)

	var e *InternalError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsInternal(err)
	assert.False(t, ok)

	ok = errors.Is(err, &InternalError{})
	assert.False(t, ok)
}

func TestFindWrappedInternal(t *testing.T) {
	err := ThrowInternal(nil, "id", "msg")
	ok := IsInternal(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsInternal(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InternalError{})
	assert.True(t, ok)

	var e *InternalError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestInternalWithRootCause(t *testing.T) {
	err := ThrowInternal(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsInternal(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InternalError{})
	assert.True(t, ok)

	var asDomain *InternalError
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

func TestWrappingAnotherInternalError(t *testing.T) {
	err := ThrowInternal(nil, "id1", "msg")
	ok := IsInternal(err)
	assert.True(t, ok)

	err = ThrowInternal(err, "id2", "msg")
	ok = IsInternal(err)
	assert.True(t, ok)

	ok = errors.Is(err, &InternalError{})
	assert.True(t, ok)

	var e *InternalError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
