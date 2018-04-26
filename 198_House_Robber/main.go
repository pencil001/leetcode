package main

import (
	"fmt"
)

func main() {
	// nums := []int{}
	nums := []int{2, 4, 1, 1, 3}
	n := rob2(nums)
	fmt.Println(n)
}

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	length := len(nums)
	earns := make([]int, len(nums))
	if length == 1 {
		earns[0] = nums[0]
	} else {
		earns[0] = nums[0]
		earns[1] = nums[1]
	}

	for i := 2; i < length; i++ {
		earns[i] = 0
		for j := 0; j < i-1; j++ {
			if earns[j]+nums[i] > earns[i] {
				earns[i] = earns[j] + nums[i]
			}
		}
	}

	max := earns[0]
	for _, v := range earns {
		if v > max {
			max = v
		}
	}
	return max
}

func rob2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
