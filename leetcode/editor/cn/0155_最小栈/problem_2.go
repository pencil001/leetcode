package p0155

import "math"

// 这里不使用额外的数组
type MinStack2 struct {
	arr []int
	min int
}

func Constructor2() MinStack2 {
	return MinStack2{
		arr: []int{},
		min: math.MaxInt64,
	}
}

func (this *MinStack2) Push(x int) {
	diff := 0
	if this.min == math.MaxInt64 {
		this.min = x
	} else {
		diff = x - this.min
	}

	this.arr = append(this.arr, diff)
	if diff < 0 {
		this.min = x
	}
}

func (this *MinStack2) Pop() {
	diff := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	if len(this.arr) == 0 {
		this.min = math.MaxInt64
	}
	if diff < 0 {
		this.min = this.min - diff
	}
}

func (this *MinStack2) Top() int {
	diff := this.arr[len(this.arr)-1]
	if diff >= 0 {
		return diff + this.min
	}

	return this.min
}

func (this *MinStack2) GetMin() int {
	return this.min
}
