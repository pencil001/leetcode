package p0150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0150(t *testing.T) {
	assert.EqualValues(t, 9, evalRPN([]string{"2", "1", "+", "3", "*"}))
	assert.EqualValues(t, 6, evalRPN([]string{"4", "13", "5", "/", "+"}))
	assert.EqualValues(t, 22, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
