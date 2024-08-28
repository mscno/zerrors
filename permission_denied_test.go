package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPermissionDeniedOnNil(t *testing.T) {
	err := ToPermissionDenied(nil, "id")
	ok := IsPermissionDenied(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToPermissionDenied(t *testing.T) {
	err := ToPermissionDenied(errors.New("not domain"), "id")
	assert.True(t, IsPermissionDenied(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrPermissionDenied, e.Code())
}

func TestIsPermissionDeniedOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsPermissionDenied(err)
	assert.False(t, ok)

}

func TestThrowPermissionDenied(t *testing.T) {
	err := ThrowPermissionDenied("create", "user", "sam")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestPermissionDeniedf(t *testing.T) {
	err := PermissionDeniedf("something happened: %s", "reason")
	ok := IsPermissionDenied(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}
