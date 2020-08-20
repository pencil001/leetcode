package p0236

import (
	"testing"

	. "github.com/pencil001/leetcode/utils"
	"github.com/stretchr/testify/assert"
)

func Test_0236(t *testing.T) {
	root := DeserializesTree("[3,5,1,6,2,0,8,null,null,7,4]")
	node := lowestCommonAncestor(root, GetTreeNodeByValue(root, 5), GetTreeNodeByValue(root, 1))
	assert.EqualValues(t, 3, node.Val)
}
