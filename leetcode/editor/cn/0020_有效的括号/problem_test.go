package p0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0020(t *testing.T) {
	assert.EqualValues(t, true, isValid("()"))
	assert.EqualValues(t, true, isValid("()[]{}"))
	assert.EqualValues(t, false, isValid("(]"))
	assert.EqualValues(t, false, isValid("([)]"))
	assert.EqualValues(t, true, isValid("{[]}"))
	assert.EqualValues(t, false, isValid("]"))
}
