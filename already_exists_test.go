package zerrors

import (
	"errors"
	"github.com/samber/oops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAlreadyExistsOnNil(t *testing.T) {
	err := ToAlreadyExists(nil, "id")
	ok := IsAlreadyExists(err)
	assert.False(t, ok)
	assert.Equal(t, err, nil)

}

func TestToAlreadyExists(t *testing.T) {
	err := ToAlreadyExists(errors.New("not domain"), "id")
	assert.True(t, IsAlreadyExists(err))

	var e oops.OopsError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrAlreadyExists, e.Code())
}

func TestIsAlreadyExistsOnNonDomain(t *testing.T) {
	err := errors.New("I am found!")
	ok := IsAlreadyExists(err)
	assert.False(t, ok)

}

func TestThrowAlreadyExists(t *testing.T) {
	err := ThrowAlreadyExists("create", "user", "sam")
	ok := IsAlreadyExists(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "cannot create 'sam' of kind 'user'")
}

func TestAlreadyExistsf(t *testing.T) {
	err := AlreadyExistsf("something happened: %s", "reason")
	ok := IsAlreadyExists(err)
	assert.True(t, ok)

	assert.Equal(t, err.Error(), "something happened: reason")
}

func TestToAlreadyExistsf(t *testing.T) {
	err := ToAlreadyExistsf(errors.New("something happened"), "key: %s", "value")
	ok := IsAlreadyExists(err)
	assert.True(t, ok)

	var e oops.OopsError
	ok = errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, ErrAlreadyExists, e.Code())

	assert.Equal(t, err.Error(), "key: value: something happened")
}
