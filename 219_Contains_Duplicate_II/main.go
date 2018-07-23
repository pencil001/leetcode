package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int, 0)
	for i, n := range nums {
		j, ok := hash[n]
		if !ok {
			hash[n] = i
		} else {
			if i-j <= k {
				return true
			} else {
				hash[n] = i
			}
		}
	}
	return false
}
