package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 3}
	i := majorityElement2(nums)
	fmt.Println(i)
}

func majorityElement(nums []int) int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}

	mk := 0
	mv := 0
	for k, v := range hash {
		if v > mv {
			mk, mv = k, v
		}
	}
	return mk
}

func majorityElement2(nums []int) int {
	count := 0
	var candidate int
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}

		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
