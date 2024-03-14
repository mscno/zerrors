package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlreadyExistsError(t *testing.T) {
	var alreadyExistsError interface{} = new(AlreadyExistsError)
	_, ok := alreadyExistsError.(AlreadyExists)
	assert.True(t, ok)
}

func TestThrowAlreadyExistsf(t *testing.T) {
	err := ThrowAlreadyExistsf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*AlreadyExistsError)
	assert.True(t, ok)
}

func TestIsErrorAlreadyExists(t *testing.T) {
	err := ThrowAlreadyExists(nil, "id", "msg")
	ok := IsErrorAlreadyExists(err)
	assert.True(t, ok)

	err = errors.New("Already Exists!")
	ok = IsErrorAlreadyExists(err)
	assert.False(t, ok)
}
