package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToResourceExhaustedOnNil(t *testing.T) {
	err := ToResourceExhausted(nil, "id")
	ok := IsResourceExhausted(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToResourceExhausted(t *testing.T) {
	err := ToResourceExhausted(errors.New("not domain"), "id")
	assert.True(t, IsResourceExhausted(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrResourceExhausted, e.Code())
}

func TestIsResourceExhaustedOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsResourceExhausted(err)
	assert.False(t, ok)

}

func TestThrowResourceExhausted(t *testing.T) {
	err := ThrowResourceExhausted("create", "user", "sam")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestResourceExhaustedf(t *testing.T) {
	err := ResourceExhaustedf("something happened: %s", "reason")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}

func TestToResourceExhaustedf(t *testing.T) {
	err := ToResourceExhaustedf(errors.New("something happened"), "key: %s", "value")
	ok := IsResourceExhausted(err)
	assert.True(t, ok)

	var e oops.OopsError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrResourceExhausted, e.Code())

	assert.Equal(t, err.Error(), "key: value: something happened")
}
