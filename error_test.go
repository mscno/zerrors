package zerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	err := errors.New("hello world")
	world := Contains(err, "hello")
	assert.True(t, world)

	mars := Contains(err, "mars")
	assert.False(t, mars)
}
