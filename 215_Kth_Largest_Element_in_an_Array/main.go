package main

import (
	"fmt"
)

func main() {
	// nums := []int{3, 2, 1, 5, 6, 4}
	// k := 2
	// nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	// k := 4
	nums := []int{2, 1}
	k := 1
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	i := quicksort(nums, 0, n-1, n-k)
	return nums[i]
}

func quicksort(nums []int, low, high, k int) int {
	if low < high {
		idx := partition(nums, low, high)
		if idx == k {
			return idx
		}

		if idx > k {
			return quicksort(nums, low, idx-1, k)
		} else {
			return quicksort(nums, idx+1, high, k)
		}
	}
	return low
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
