package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	n := maxSubArray(nums)
	fmt.Println(n)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sums := make([]int, len(nums))
	for i, v := range nums {
		sums[i] = v
		if i-1 >= 0 {
			if sums[i] < sums[i-1]+v {
				sums[i] = sums[i-1] + v
			}
		}
	}

	max := sums[0]
	for _, v := range sums {
		if v > max {
			max = v
		}
	}
	return max
}
