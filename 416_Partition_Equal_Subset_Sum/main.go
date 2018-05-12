package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 5, 11, 5} //true
	// nums := []int{1, 2, 3, 5} //false
	// nums := []int{1, 2, 5} //false
	nums := []int{66, 90, 7, 6, 32, 16, 2, 78, 69, 88, 85, 26, 3, 9, 58, 65, 30, 96, 11, 31, 99, 49, 63, 83, 79, 97, 20, 64, 81, 80, 25, 69, 9, 75, 23, 70, 26, 71, 25, 54, 1, 40, 41, 82, 32, 10, 26, 33, 50, 71, 5, 91, 59, 96, 9, 15, 46, 70, 26, 32, 49, 35, 80, 21, 34, 95, 51, 66, 17, 71, 28, 88, 46, 21, 31, 71, 42, 2, 98, 96, 40, 65, 92, 43, 68, 14, 98, 38, 13, 77, 14, 13, 60, 79, 52, 46, 9, 13, 25, 8}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}

	return subsetSum(nums, sum/2)
}

func subsetSum(nums []int, s int) bool {
	dp := make([]bool, s+1)
	dp[0] = true
	for _, n := range nums {
		for i := s; i >= n; i-- {
			dp[i] = dp[i] || dp[i-n]
		}
	}
	return dp[s]
}
