package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	i := singleNumber(nums)
	fmt.Print(i)
}

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}
