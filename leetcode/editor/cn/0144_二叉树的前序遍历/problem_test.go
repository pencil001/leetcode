package p0144

import (
	"testing"

	. "github.com/pencil001/leetcode/utils"
	"github.com/stretchr/testify/assert"
)

func Test_0144(t *testing.T) {
	assert.EqualValues(t, []int{1, 2, 3}, preorderTraversal(DeserializesTree("[1,null,2,3]")))
}
