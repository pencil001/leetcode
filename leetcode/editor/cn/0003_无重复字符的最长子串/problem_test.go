package p0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0003(t *testing.T) {
	assert.EqualValues(t, 0, lengthOfLongestSubstring(""))
	assert.EqualValues(t, 1, lengthOfLongestSubstring(" "))
	assert.EqualValues(t, 2, lengthOfLongestSubstring("au"))
	assert.EqualValues(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.EqualValues(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.EqualValues(t, 3, lengthOfLongestSubstring("pwwkew"))
}
