package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}

type MinStack struct {
	normals []int
	np      int
	mins    []int
	mp      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		normals: make([]int, 10000),
		np:      -1,
		mins:    make([]int, 10000),
		mp:      -1,
	}
}

func (this *MinStack) Push(x int) {
	this.np++
	this.normals[this.np] = x

	var min int
	if this.mp == -1 {
		min = x
	} else {
		min = this.mins[this.mp]
		if x < min {
			min = x
		}
	}
	this.mp++
	this.mins[this.mp] = min
}

func (this *MinStack) Pop() {
	this.np--
	this.mp--
}

func (this *MinStack) Top() int {
	if this.np == -1 {
		return 0
	} else {
		return this.normals[this.np]
	}
}

func (this *MinStack) GetMin() int {
	if this.mp == -1 {
		return 0
	} else {
		return this.mins[this.mp]
	}
}
