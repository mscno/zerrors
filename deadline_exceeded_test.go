package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDeadlineExceededOnNil(t *testing.T) {
	err := ToDeadlineExceeded(nil, "id")
	ok := IsDeadlineExceeded(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToDeadlineExceeded(t *testing.T) {
	err := ToDeadlineExceeded(errors.New("not domain"), "id")
	assert.True(t, IsDeadlineExceeded(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrDeadlineExceeded, e.Code())
}

func TestIsDeadlineExceededOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsDeadlineExceeded(err)
	assert.False(t, ok)

}

func TestThrowDeadlineExceeded(t *testing.T) {
	err := ThrowDeadlineExceeded("create", "user", "sam")
	ok := IsDeadlineExceeded(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestDeadlineExceededf(t *testing.T) {
	err := DeadlineExceededf("something happened: %s", "reason")
	ok := IsDeadlineExceeded(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}

func TestToDeadlineExceededf(t *testing.T) {
	err := ToDeadlineExceededf(errors.New("something happened"), "key: %s", "value")
	ok := IsDeadlineExceeded(err)
	assert.True(t, ok)

	var e oops.OopsError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrDeadlineExceeded, e.Code())

	assert.Equal(t, err.Error(), "key: value: something happened")
}
