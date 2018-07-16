package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS2(nums))
}

func lengthOfLIS(nums []int) int {
	lens := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		lens[i] = 1
	}

	max := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if lens[j]+1 > lens[i] {
					lens[i] = lens[j] + 1
				}
			}
		}
		if lens[i] > max {
			max = lens[i]
		}
	}

	return max
}

func lengthOfLIS2(nums []int) int {
	dp := []int{}

	for _, n := range nums {
		i := sort.SearchInts(dp, n)
		if i < len(dp) {
			dp[i] = n
		} else {
			dp = append(dp, n)
		}
	}
	return len(dp)
}
