package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 1, 1, 1, 1}
	// S := 3 // 5
	// nums := []int{5, 2, 7, 3, 9}
	// S := 16 // 2
	// nums := []int{1}
	// S := 1 // 1
	// nums := []int{1, 0}
	// S := 1 // 2
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	S := 1 // 256
	fmt.Println(findTargetSumWays(nums, S))
}

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum < S || (sum+S)%2 == 1 {
		return 0
	}

	return subsetSum(nums, (sum+S)>>1)
}

func subsetSum(nums []int, s int) int {
	dp := make([]int, s+1)
	dp[0] = 1
	for _, n := range nums {
		for i := s; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[s]
}
