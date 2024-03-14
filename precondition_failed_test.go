package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreconditionFailedError(t *testing.T) {
	var err interface{} = new(PreconditionFailedError)
	_, ok := err.(PreconditionFailed)
	assert.True(t, ok)
}

func TestThrowPreconditionFailedf(t *testing.T) {
	err := ThrowPreconditionFailedf(nil, "id", "msg")
	//nolint:errorlint
	_, ok := err.(*PreconditionFailedError)
	assert.True(t, ok)
}

func TestIsPreconditionFailed(t *testing.T) {
	err := ThrowPreconditionFailed(nil, "id", "msg")
	ok := IsPreconditionFailed(err)
	assert.True(t, ok)

	err = errors.New("Precondition failed!")
	ok = IsPreconditionFailed(err)
	assert.False(t, ok)
}
