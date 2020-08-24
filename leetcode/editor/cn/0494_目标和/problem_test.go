package p0494

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0494(t *testing.T) {
	assert.EqualValues(t, 5, findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
