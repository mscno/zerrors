package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUnknownOnNil(t *testing.T) {
	err := ToUnknown(nil, "id")
	ok := IsUnknown(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToUnknown(t *testing.T) {
	err := ToUnknown(errors.New("not domain"), "id")
	assert.True(t, IsUnknown(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrUnknown, e.Code())
}

func TestIsUnknownOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsUnknown(err)
	assert.False(t, ok)

}

func TestThrowUnknown(t *testing.T) {
	err := ThrowUnknown("create", "user", "sam")
	ok := IsUnknown(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestUnknownf(t *testing.T) {
	err := Unknownf("something happened: %s", "reason")
	ok := IsUnknown(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}
