package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToNotFoundOnNil(t *testing.T) {
	err := ToNotFound(nil, "id")
	ok := IsNotFound(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToNotFound(t *testing.T) {
	err := ToNotFound(errors.New("not domain"), "id")
	assert.True(t, IsNotFound(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrNotFound, e.Code())
}

func TestIsNotFoundOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsNotFound(err)
	assert.False(t, ok)

}

func TestThrowNotFound(t *testing.T) {
	err := ThrowNotFound("create", "user", "sam")
	ok := IsNotFound(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestNotFoundf(t *testing.T) {
	err := NotFoundf("something happened: %s", "reason")
	ok := IsNotFound(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}
