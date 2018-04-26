package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	// nums := []int{3, 2, 4}
	// target := 6
	pos := twoSum(nums, target)
	fmt.Println(pos)
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		j, ok := hash[target-v]
		if ok {
			return []int{j, i}
		}

		hash[v] = i
	}
	return []int{}
}
