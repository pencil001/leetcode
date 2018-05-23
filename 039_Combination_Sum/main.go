package main

import (
	"fmt"
	"sort"
)

func main() {
	// candidates := []int{2, 3, 6, 7}
	// target := 7
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	cands := []int{}
	nums := []int{}
	for _, v := range candidates {
		n := target / v
		if n > 0 {
			cands = append(cands, v)
			nums = append(nums, n)
		}
	}

	res := [][]int{}
	backtrack(cands, nums, target, []int{}, &res)
	return res
}

func backtrack(cands, nums []int, target int, cur []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i := 0; i < len(cands); i++ {
		v := cands[i]
		n := nums[i]

		numsCopy := []int{}
		l := len(cur)
		cur = append(cur, v)
		if n-1 == 0 {
			numsCopy = append(numsCopy, nums[i+1:]...)
			backtrack(cands[i+1:], numsCopy, target-v, cur, res)
		} else {
			numsCopy = append(numsCopy, nums[i:]...)
			numsCopy[0]--
			backtrack(cands[i:], numsCopy, target-v, cur, res)
		}
		cur = cur[:l]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	backtrack2(candidates, target, 0, []int{}, &res)
	return res
}

func backtrack2(candidates []int, target int, idx int, cur []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i := idx; i < len(candidates); i++ {
		v := candidates[i]
		l := len(cur)
		cur = append(cur, v)

		backtrack2(candidates, target-v, i, cur, res)

		cur = cur[:l]
	}
}
