package p0155

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0155(t *testing.T) {
	assert.EqualValues(t,
		"[null,null,null,null,-3,null,0,-2]",
		processInput(
			`["MinStack","push","push","push","getMin","pop","top","getMin"]`,
			`[[],[-2],[0],[-3],[],[],[],[]]`,
		))

	assert.EqualValues(t,
		"[null,null,null,null,2147483647,null,2147483646,null,2147483646,null,null,2147483647,2147483647,null,-2147483648,-2147483648,null,2147483647]",
		processInput(
			`["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]`,
			`[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]`,
		))
}

func processInput(strCmd, strArg string) string {
	commands := make([]string, 0)
	_ = json.Unmarshal([]byte(strCmd), &commands)

	args := make([][]int, 0)
	_ = json.Unmarshal([]byte(strArg), &args)

	var results []interface{}
	var obj MinStack2
	for i, c := range commands {
		switch c {
		case "MinStack":
			obj = Constructor2()
			results = append(results, nil)
		case "push":
			obj.Push(args[i][0])
			results = append(results, nil)
		case "pop":
			obj.Pop()
			results = append(results, nil)
		case "top":
			results = append(results, obj.Top())
		case "getMin":
			results = append(results, obj.GetMin())
		}
	}

	bs, _ := json.Marshal(results)
	return string(bs)
}
