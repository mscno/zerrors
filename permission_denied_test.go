package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionDeniedError(t *testing.T) {
	var err interface{}
	err = new(PermissionDeniedError)
	_, ok := err.(PermissionDenied)
	assert.True(t, ok)
}

func TestToPermissionDeniedf(t *testing.T) {
	err := ToPermissionDeniedf(nil, "id", "msg")
	_, ok := err.(*PermissionDeniedError)
	assert.True(t, ok)
}

func TestIsPermissionDenied(t *testing.T) {
	err := ToPermissionDenied(nil, "id", "msg")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	ok = errors.Is(err, &PermissionDeniedError{})
	assert.True(t, ok)

	var e *PermissionDeniedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsPermissionDenied(err)
	assert.False(t, ok)

	ok = errors.Is(err, &PermissionDeniedError{})
	assert.False(t, ok)
}

func TestFindWrappedPermissionDenied(t *testing.T) {
	err := ToPermissionDenied(nil, "id", "msg")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsPermissionDenied(err)
	assert.True(t, ok)

	ok = errors.Is(err, &PermissionDeniedError{})
	assert.True(t, ok)

	var e *PermissionDeniedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestPermissionDeniedWithRootCause(t *testing.T) {
	err := ToPermissionDenied(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	ok = errors.Is(err, &PermissionDeniedError{})
	assert.True(t, ok)

	var asDomain *PermissionDeniedError
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

func TestWrappingAnotherPermissionDeniedError(t *testing.T) {
	err := ToPermissionDenied(nil, "id1", "msg")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	err = ToPermissionDenied(err, "id2", "msg")
	ok = IsPermissionDenied(err)
	assert.True(t, ok)

	ok = errors.Is(err, &PermissionDeniedError{})
	assert.True(t, ok)

	var e *PermissionDeniedError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
