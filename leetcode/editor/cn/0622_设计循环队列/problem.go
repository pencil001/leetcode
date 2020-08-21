package p0622

//leetcode submit region begin(Prohibit modification and deletion)
type MyCircularQueue struct {
	head int
	size int
	arr  []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	arr := make([]int, k, k)
	return MyCircularQueue{head: 0, size: 0, arr: arr}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	curTail := (this.head + this.size - 1) % len(this.arr)
	nextTail := (curTail + 1) % len(this.arr)
	this.arr[nextTail] = value
	this.size++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % len(this.arr)
	this.size--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.arr[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	tail := (this.head + this.size - 1) % len(this.arr)
	return this.arr[tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.arr)
}

//leetcode submit region end(Prohibit modification and deletion)
