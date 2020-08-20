package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

func NewTree(nums []interface{}) *TreeNode {
	var newTreeEx func([]interface{}, int) *TreeNode
	newTreeEx = func(nums []interface{}, rootIdx int) *TreeNode {
		if rootIdx >= len(nums) {
			return nil
		}

		v := nums[rootIdx]
		if v == nil {
			return nil
		}

		n := &TreeNode{Val: v.(int)}
		n.Left = newTreeEx(nums, 2*rootIdx+1)
		n.Right = newTreeEx(nums, 2*rootIdx+2)
		return n
	}

	return newTreeEx(nums, 0)
}

func WalkTree(t *TreeNode) {
	if t == nil {
		return
	}

	WalkTree(t.Left)
	fmt.Printf("%v;", t.Val)
	WalkTree(t.Right)
}

// SerializesTree a tree to a single string.
func SerializesTree(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	bfsQ := []*TreeNode{}

	q := []*TreeNode{root}
	var node *TreeNode
	for len(q) > 0 {
		node, q = q[0], q[1:]
		bfsQ = append(bfsQ, node)
		if node != nil {
			q = append(q, node.Left, node.Right)
		}
	}

	last := len(bfsQ)
	for i := last - 1; i >= 0; i-- {
		if bfsQ[i] == nil {
			last--
		} else {
			break
		}
	}
	bfsQ = bfsQ[:last]

	sb := strings.Builder{}
	sb.WriteString("[")

	for i := 0; i < len(bfsQ); i++ {
		if sb.Len() != 1 {
			sb.WriteString(",")
		}

		if bfsQ[i] != nil {
			sb.WriteString(fmt.Sprint(bfsQ[i].Val))
		} else {
			sb.WriteString("null")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

// DeserializesTree your encoded data to tree.
func DeserializesTree(data string) *TreeNode {
	arr := []interface{}{}

	if err := json.Unmarshal([]byte(data), &arr); err != nil {
		panic(err)
	}

	cnt := len(arr)
	if cnt == 0 {
		return nil
	}

	i := 0
	root := &TreeNode{Val: int(arr[i].(float64))}
	preLevel := []*TreeNode{root}
	i++

	isLeft := true
	for len(preLevel) > 0 {
		tmpL := []*TreeNode{}

		for j := 0; j < len(preLevel); {
			parent := preLevel[j]
			if parent == nil {
				j++
				continue
			}

			if i >= len(arr) {
				break
			}

			var node *TreeNode
			if arr[i] != nil {
				node = &TreeNode{Val: int(arr[i].(float64))}
			}
			tmpL = append(tmpL, node)
			i++

			if isLeft {
				parent.Left = node
				isLeft = false
			} else {
				parent.Right = node
				isLeft = true
				j++
			}
		}

		preLevel = tmpL
	}

	return root
}

func GetTreeNodeByValue(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == v {
		return root
	}

	l := GetTreeNodeByValue(root.Left, v)
	r := GetTreeNodeByValue(root.Right, v)
	if l == nil {
		return r
	}
	return l
}
