package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInternalOnNil(t *testing.T) {
	err := ToInternal(nil, "id")
	ok := IsInternal(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToInternal(t *testing.T) {
	err := ToInternal(errors.New("not domain"), "id")
	assert.True(t, IsInternal(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrInternal, e.Code())
}

func TestIsInternalOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsInternal(err)
	assert.False(t, ok)

}

func TestThrowInternal(t *testing.T) {
	err := ThrowInternal("create", "user", "sam")
	ok := IsInternal(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestInternalf(t *testing.T) {
	err := Internalf("something happened: %s", "reason")
	ok := IsInternal(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}

func TestToInternalf(t *testing.T) {
	err := ToInternalf(errors.New("something happened"), "key: %s", "value")
	ok := IsInternal(err)
	assert.True(t, ok)

	var e oops.OopsError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrInternal, e.Code())

	assert.Equal(t, err.Error(), "key: value: something happened")
}
