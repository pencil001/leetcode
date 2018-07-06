package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0)) // 0
}

func searchInsert(nums []int, target int) int {
	if nums == nil {
		return 0
	}

	begin := 0
	end := len(nums) - 1

	for begin <= end {
		mid := (begin + end) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	return begin
}
