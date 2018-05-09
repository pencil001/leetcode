package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	for _, v := range res {
		fmt.Println(v)
	}
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	sort.Ints(nums)
	backtrack(&res, []int{}, nums)
	return res
}

func backtrack(res *[][]int, cur []int, nums []int) {
	if len(nums) == 0 {
		return
	}

	for i := 0; i < len(nums); i++ {
		temp := append(cur, nums[i])
		*res = append(*res, append([]int{}, temp...))
		backtrack(res, temp, nums[i+1:])
	}
}
