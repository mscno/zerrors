package zerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMethod(t *testing.T) {
	err := ThrowError(nil, "id", "msg")
	expected := "ID=id Message=msg"
	assert.Equal(t, expected, err.Error())

	err = ThrowError(err, "subID", "subMsg")
	subExptected := "ID=subID Message=subMsg Parent=(ID=id Message=msg)"
	assert.Equal(t, subExptected, err.Error())
}
