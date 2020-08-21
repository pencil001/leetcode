package p0279

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0279(t *testing.T) {
	assert.EqualValues(t, 1, numSquares(4))
	assert.EqualValues(t, 1, numSquares(1))
	assert.EqualValues(t, 3, numSquares(12))
	assert.EqualValues(t, 2, numSquares(13))
}
