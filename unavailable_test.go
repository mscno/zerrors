package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUnavailableOnNil(t *testing.T) {
	err := ToUnavailable(nil, "id")
	ok := IsUnavailable(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToUnavailable(t *testing.T) {
	err := ToUnavailable(errors.New("not domain"), "id")
	assert.True(t, IsUnavailable(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrUnavailable, e.Code())
}

func TestIsUnavailableOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsUnavailable(err)
	assert.False(t, ok)

}

func TestThrowUnavailable(t *testing.T) {
	err := ThrowUnavailable("create", "user", "sam")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestUnavailablef(t *testing.T) {
	err := Unavailablef("something happened: %s", "reason")
	ok := IsUnavailable(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}
