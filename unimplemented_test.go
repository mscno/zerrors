package zerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnimplementedError(t *testing.T) {
	var err interface{}
	err = new(UnimplementedError)
	_, ok := err.(Unimplemented)
	assert.True(t, ok)
}

func TestToUnimplementedf(t *testing.T) {
	err := ToUnimplementedf(nil, "id", "msg")
	_, ok := err.(*UnimplementedError)
	assert.True(t, ok)
}

func TestIsUnimplemented(t *testing.T) {
	err := ToUnimplemented(nil, "id", "msg")
	ok := IsUnimplemented(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnimplementedError{})
	assert.True(t, ok)

	var e *UnimplementedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

	err = errors.New("I am found!")
	ok = IsUnimplemented(err)
	assert.False(t, ok)

	ok = errors.Is(err, &UnimplementedError{})
	assert.False(t, ok)
}

func TestFindWrappedUnimplemented(t *testing.T) {
	err := ToUnimplemented(nil, "id", "msg")
	ok := IsUnimplemented(err)
	assert.True(t, ok)

	err = fmt.Errorf("wrapped %w", err)
	ok = IsUnimplemented(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnimplementedError{})
	assert.True(t, ok)

	var e *UnimplementedError
	ok = errors.As(err, &e)
	assert.True(t, ok)

}

func TestUnimplementedWithRootCause(t *testing.T) {
	err := ToUnimplemented(fmt.Errorf("not domain"), "id", "no wrap message")
	ok := IsUnimplemented(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnimplementedError{})
	assert.True(t, ok)

	var asDomain *UnimplementedError
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

func TestWrappingAnotherUnimplementedError(t *testing.T) {
	err := ToUnimplemented(nil, "id1", "msg")
	ok := IsUnimplemented(err)
	assert.True(t, ok)

	err = ToUnimplemented(err, "id2", "msg")
	ok = IsUnimplemented(err)
	assert.True(t, ok)

	ok = errors.Is(err, &UnimplementedError{})
	assert.True(t, ok)

	var e *UnimplementedError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, "id2", e.Zerror.ID)

	msg := "ID=id2 Message=msg Parent=(ID=id1 Message=msg)"
	assert.Equal(t, msg, err.Error())
}
