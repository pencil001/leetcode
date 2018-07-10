package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	// nums := []int{1, 2, 0}
	sortColors2(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	counts := []int{0, 0, 0}
	for _, n := range nums {
		counts[n]++
	}

	k := 0
	for i := 0; i < 3; i++ {
		n := counts[i]
		for j := k; j < k+n; j++ {
			nums[j] = i
		}
		k = k + n
	}
}

func sortColors2(nums []int) {
	b := 0
	e := len(nums) - 1
	c := 0
	for c <= e {
		n := nums[c]
		if n == 0 {
			nums[b], nums[c] = nums[c], nums[b]
			c++
			b++
		} else if n == 2 {
			nums[e], nums[c] = nums[c], nums[e]
			e--
		} else {
			c++
		}
	}
}
