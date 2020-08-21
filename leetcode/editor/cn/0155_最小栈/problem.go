package p0155

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	arr  []int
	ap   int
	mins []int
	mp   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		arr:  make([]int, 10000),
		ap:   -1,
		mins: make([]int, 10000),
		mp:   -1,
	}
}

func (this *MinStack) Push(x int) {
	this.ap++
	this.arr[this.ap] = x

	var min int
	if this.mp == -1 {
		min = x
	} else {
		min = this.GetMin()
	}

	if x < min {
		min = x
	}

	this.mp++
	this.mins[this.mp] = min
}

func (this *MinStack) Pop() {
	this.ap--
	this.mp--
}

func (this *MinStack) Top() int {
	return this.arr[this.ap]
}

func (this *MinStack) GetMin() int {
	return this.mins[this.mp]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
