package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 3}
	nums := []int{5, 4, 6, 2}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	backtrack(&res, []int{}, nums)
	return res
}

func backtrack(res *[][]int, cur []int, nums []int) {
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(cur, nums[i]) {
			continue
		}

		cur = append(cur, nums[i])
		backtrack(res, cur, nums)
		cur = cur[:len(cur)-1]
	}
}

func contains(cur []int, v int) bool {
	for i := 0; i < len(cur); i++ {
		if cur[i] == v {
			return true
		}
	}
	return false
}
