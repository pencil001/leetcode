package p0076

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0076(t *testing.T) {
	assert.EqualValues(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.EqualValues(t, "a", minWindow("a", "a"))
}
