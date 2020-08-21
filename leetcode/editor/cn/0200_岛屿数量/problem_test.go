package p0200

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0200(t *testing.T) {
	//输入:
	//[
	//  ['1','1','1','1','0'],
	//  ['1','1','0','1','0'],
	//  ['1','1','0','0','0'],
	//  ['0','0','0','0','0']
	//]
	//输出: 1
	assert.EqualValues(t, 1, processInput(
		`[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]`))

	//输入:
	//[
	//  ['1','1','0','0','0'],
	//  ['1','1','0','0','0'],
	//  ['0','0','1','0','0'],
	//  ['0','0','0','1','1']
	//]
	//输出: 3
	assert.EqualValues(t, 3, processInput(
		`[["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]`))
}

func processInput(strArg string) int {
	args := make([][]string, 0)
	if err := json.Unmarshal([]byte(strArg), &args); err != nil {
		panic(err)
	}

	bytes := make([][]byte, len(args))
	for i, as := range args {
		bs := make([]byte, len(as))
		for j, a := range as {
			bs[j] = []byte(a)[0]
		}
		bytes[i] = bs
	}

	index := 2
	numIslands := []func([][]byte) int{
		numIslandsDfs1,
		numIslandsDfs2,
		numIslandsBfs,
	}

	return numIslands[index](bytes)
}
