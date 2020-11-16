package p0567

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0567(t *testing.T) {
	assert.EqualValues(t, true, checkInclusion("ab", "eidbaooo"))
	assert.EqualValues(t, false, checkInclusion("ab", "eidboaoo"))
}
