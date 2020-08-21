package p0752

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0752(t *testing.T) {
	assert.EqualValues(t, 6, openLockDoubleBfs([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	assert.EqualValues(t, 1, openLockDoubleBfs([]string{"8888"}, "0009"))
	assert.EqualValues(t, -1, openLockDoubleBfs([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
	assert.EqualValues(t, -1, openLockDoubleBfs([]string{"0000"}, "8888"))
}
