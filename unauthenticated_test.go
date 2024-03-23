package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnauthenticatedError(t *testing.T) {
	var err interface{}
	err = new(UnauthenticatedError)
	_, ok := err.(Unauthenticated)
	assert.True(t, ok)
}

func TestThrowUnauthenticatedf(t *testing.T) {
	err := ThrowUnauthenticatedf(nil, "id", "msg")
	_, ok := err.(*UnauthenticatedError)
	assert.True(t, ok)
}

func TestIsUnauthenticated(t *testing.T) {
	err := ThrowUnauthenticated(nil, "id", "msg")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnauthenticatedError{})
	assert.True(t, ok)

	var e *UnauthenticatedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnauthenticated(err)
	assert.False(t, ok)

	ok = errors.Is(err, &UnauthenticatedError{})
	assert.False(t, ok)
}

func TestFindWrappedUnauthenticated(t *testing.T) {
	err := ThrowUnauthenticated(nil, "id", "msg")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsUnauthenticated(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnauthenticatedError{})
	assert.True(t, ok)

	var e *UnauthenticatedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestUnauthenticatedWithRootCause(t *testing.T) {
	err := ThrowUnauthenticated(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnauthenticatedError{})
	assert.True(t, ok)

	var asDomain *UnauthenticatedError
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

func TestWrappingAnotherUnauthenticatedError(t *testing.T) {
	err := ThrowUnauthenticated(nil, "id1", "msg")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	err = ThrowUnauthenticated(err, "id2", "msg")
	ok = IsUnauthenticated(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnauthenticatedError{})
	assert.True(t, ok)

	var e *UnauthenticatedError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
