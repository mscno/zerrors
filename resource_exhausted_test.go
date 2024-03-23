package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceExhaustedError(t *testing.T) {
	var err interface{}
	err = new(ResourceExhaustedError)
	_, ok := err.(ResourceExhausted)
	assert.True(t, ok)
}

func TestThrowResourceExhaustedf(t *testing.T) {
	err := ThrowResourceExhaustedf(nil, "id", "msg")
	_, ok := err.(*ResourceExhaustedError)
	assert.True(t, ok)
}

func TestIsResourceExhausted(t *testing.T) {
	err := ThrowResourceExhausted(nil, "id", "msg")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	ok = errors.Is(err, &ResourceExhaustedError{})
	assert.True(t, ok)

	var e *ResourceExhaustedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsResourceExhausted(err)
	assert.False(t, ok)

	ok = errors.Is(err, &ResourceExhaustedError{})
	assert.False(t, ok)
}

func TestFindWrappedResourceExhausted(t *testing.T) {
	err := ThrowResourceExhausted(nil, "id", "msg")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsResourceExhausted(err)
	assert.True(t, ok)

	ok = errors.Is(err, &ResourceExhaustedError{})
	assert.True(t, ok)

	var e *ResourceExhaustedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestResourceExhaustedWithRootCause(t *testing.T) {
	err := ThrowResourceExhausted(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	ok = errors.Is(err, &ResourceExhaustedError{})
	assert.True(t, ok)

	var asDomain *ResourceExhaustedError
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

func TestWrappingAnotherResourceExhaustedError(t *testing.T) {
	err := ThrowResourceExhausted(nil, "id1", "msg")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	err = ThrowResourceExhausted(err, "id2", "msg")
	ok = IsResourceExhausted(err)
	assert.True(t, ok)

	ok = errors.Is(err, &ResourceExhaustedError{})
	assert.True(t, ok)

	var e *ResourceExhaustedError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
