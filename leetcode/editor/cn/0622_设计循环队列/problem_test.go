package p0622

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0622(t *testing.T) {
	assert.EqualValues(t,
		"[null,true,true,true,false,3,true,true,true,4]",
		processInput(
			`["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]`,
			`[[3],[1],[2],[3],[4],[],[],[],[4],[]]`,
		))

	assert.EqualValues(t,
		"[null,true,6,6,true,true,5,true,-1,false,false,false]",
		processInput(
			`["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear","deQueue","Front","deQueue","deQueue","deQueue"]`,
			`[[6],[6],[],[],[],[5],[],[],[],[],[],[]]`,
		))
}

func processInput(strCmd, strArg string) string {
	commands := make([]string, 0)
	_ = json.Unmarshal([]byte(strCmd), &commands)

	args := make([][]int, 0)
	_ = json.Unmarshal([]byte(strArg), &args)

	var results []interface{}
	var obj MyCircularQueue
	for i, c := range commands {
		switch c {
		case "MyCircularQueue":
			obj = Constructor(args[i][0])
			results = append(results, nil)
		case "enQueue":
			results = append(results, obj.EnQueue(args[i][0]))
		case "deQueue":
			results = append(results, obj.DeQueue())
		case "Rear":
			results = append(results, obj.Rear())
		case "Front":
			results = append(results, obj.Front())
		case "isFull":
			results = append(results, obj.IsFull())
		case "isEmpty":
			results = append(results, obj.IsEmpty())
		}
	}

	bs, _ := json.Marshal(results)
	return string(bs)
}
