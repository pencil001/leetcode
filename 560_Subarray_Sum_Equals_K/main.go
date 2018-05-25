package main

import "fmt"

func main() {
	// nums := []int{1, 1, 1}
	// k := 2
	nums := []int{1, 1, 1, 2}
	k := 2
	fmt.Println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += sumAll(nums[i:], k)
	}
	return res
}

func sumAll(nums []int, k int) int {
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == k {
			res++
		}
	}
	return res
}
