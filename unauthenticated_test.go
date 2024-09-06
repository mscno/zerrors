package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUnauthenticatedOnNil(t *testing.T) {
	err := ToUnauthenticated(nil, "id")
	ok := IsUnauthenticated(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToUnauthenticated(t *testing.T) {
	err := ToUnauthenticated(errors.New("not domain"), "id")
	assert.True(t, IsUnauthenticated(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrUnauthenticated, e.Code())
}

func TestIsUnauthenticatedOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsUnauthenticated(err)
	assert.False(t, ok)

}

func TestThrowUnauthenticated(t *testing.T) {
	err := ThrowUnauthenticated("create", "user", "sam")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestUnauthenticatedf(t *testing.T) {
	err := Unauthenticatedf("something happened: %s", "reason")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}

func TestToUnauthenticatedf(t *testing.T) {
	err := ToUnauthenticatedf(errors.New("something happened"), "key: %s", "value")
	ok := IsUnauthenticated(err)
	assert.True(t, ok)

	var e oops.OopsError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrUnauthenticated, e.Code())

	assert.Equal(t, err.Error(), "key: value: something happened")
}
