package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFailedPreconditionOnNil(t *testing.T) {
	err := ToFailedPrecondition(nil, "id")
	ok := IsFailedPrecondition(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToFailedPrecondition(t *testing.T) {
	err := ToFailedPrecondition(errors.New("not domain"), "id")
	assert.True(t, IsFailedPrecondition(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrFailedPrecondition, e.Code())
}

func TestIsFailedPreconditionOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsFailedPrecondition(err)
	assert.False(t, ok)

}

func TestThrowFailedPrecondition(t *testing.T) {
	err := ThrowFailedPrecondition("create", "user", "sam")
	ok := IsFailedPrecondition(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestFailedPreconditionf(t *testing.T) {
	err := FailedPreconditionf("something happened: %s", "reason")
	ok := IsFailedPrecondition(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}
