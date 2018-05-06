package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf2(nums))
}

func productExceptSelf(nums []int) []int {
	forward := make([]int, len(nums))
	backward := make([]int, len(nums))

	for i := 0; i < len(nums)-1; i++ {
		if i-1 >= 0 {
			forward[i] = forward[i-1] * nums[i]
		} else {
			forward[i] = nums[i]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i+1 < len(nums) {
			backward[i] = backward[i+1] * nums[i]
		} else {
			backward[i] = nums[i]
		}
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		left := 1
		if i-1 >= 0 {
			left = forward[i-1]
		}
		right := 1
		if i+1 < len(nums) {
			right = backward[i+1]
		}
		res[i] = left * right
	}
	return res
}

func productExceptSelf2(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	acc := 1
	for i := 0; i < len(nums)-1; i++ {
		acc *= nums[i]
		res[i+1] = acc
	}

	acc = 1
	for i := len(nums) - 1; i > 0; i-- {
		acc *= nums[i]
		res[i-1] *= acc
	}
	return res
}
