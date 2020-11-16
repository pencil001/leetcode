package p0438

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0438(t *testing.T) {
	assert.EqualValues(t, []int{0, 6}, findAnagrams("cbaebabacd", "abc"))
	assert.EqualValues(t, []int{0, 1, 2}, findAnagrams("abab", "ab"))
}
