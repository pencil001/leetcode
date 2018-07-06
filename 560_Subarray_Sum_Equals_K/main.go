package main

import "fmt"

func main() {
	fmt.Println(subarraySum2([]int{1, 1, 1}, 2))                      // should be 2
	fmt.Println(subarraySum2([]int{-1, -1, 1}, 1))                    // should be 1
	fmt.Println(subarraySum2([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0)) // should be 55
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

func subarraySum2(nums []int, k int) int {
	hash := make(map[int]int)
	hash[0] = 1

	sum := 0
	count := 0
	for _, v := range nums {
		sum += v
		if n, ok := hash[sum-k]; ok {
			count += n
		}
		hash[sum]++
	}
	return count
}
