package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUnimplementedOnNil(t *testing.T) {
	err := ToUnimplemented(nil, "id")
	ok := IsUnimplemented(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToUnimplemented(t *testing.T) {
	err := ToUnimplemented(errors.New("not domain"), "id")
	assert.True(t, IsUnimplemented(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrUnimplemented, e.Code())
}

func TestIsUnimplementedOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsUnimplemented(err)
	assert.False(t, ok)

}

func TestThrowUnimplemented(t *testing.T) {
	err := ThrowUnimplemented("create", "user", "sam")
	ok := IsUnimplemented(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}
