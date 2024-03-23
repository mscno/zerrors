package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnavailableError(t *testing.T) {
	var err interface{}
	err = new(UnavailableError)
	_, ok := err.(Unavailable)
	assert.True(t, ok)
}

func TestThrowUnavailablef(t *testing.T) {
	err := ThrowUnavailablef(nil, "id", "msg")
	_, ok := err.(*UnavailableError)
	assert.True(t, ok)
}

func TestIsUnavailable(t *testing.T) {
	err := ThrowUnavailable(nil, "id", "msg")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnavailableError{})
	assert.True(t, ok)

	var e *UnavailableError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnavailable(err)
	assert.False(t, ok)

	ok = errors.Is(err, &UnavailableError{})
	assert.False(t, ok)
}

func TestFindWrappedUnavailable(t *testing.T) {
	err := ThrowUnavailable(nil, "id", "msg")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsUnavailable(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnavailableError{})
	assert.True(t, ok)

	var e *UnavailableError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestUnavailableWithRootCause(t *testing.T) {
	err := ThrowUnavailable(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnavailableError{})
	assert.True(t, ok)

	var asDomain *UnavailableError
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

func TestWrappingAnotherUnavailableError(t *testing.T) {
	err := ThrowUnavailable(nil, "id1", "msg")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	err = ThrowUnavailable(err, "id2", "msg")
	ok = IsUnavailable(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnavailableError{})
	assert.True(t, ok)

	var e *UnavailableError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
