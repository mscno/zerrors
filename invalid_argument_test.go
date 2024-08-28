package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInvalidArgumentOnNil(t *testing.T) {
	err := ToInvalidArgument(nil, "id")
	ok := IsInvalidArgument(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToInvalidArgument(t *testing.T) {
	err := ToInvalidArgument(errors.New("not domain"), "id")
	assert.True(t, IsInvalidArgument(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrInvalidArgument, e.Code())
}

func TestIsInvalidArgumentOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsInvalidArgument(err)
	assert.False(t, ok)

}

func TestThrowInvalidArgument(t *testing.T) {
	err := ThrowInvalidArgument("create", "user", "sam")
	ok := IsInvalidArgument(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestInvalidArgumentf(t *testing.T) {
	err := InvalidArgumentf("something happened: %s", "reason")
	ok := IsInvalidArgument(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}
