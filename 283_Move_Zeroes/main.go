package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 3, 0, 5, 0, 6}
	moveZeros2(nums)
	fmt.Print(nums)
}

func moveZeros2(nums []int) {
	for lastNonZeroFoundAt, cur := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[lastNonZeroFoundAt], nums[cur] = nums[cur], 0
			lastNonZeroFoundAt++
		}
	}
}

func moveZeroes(nums []int) {
	firstZeroIdx := getZeroIdx(nums, 0)
	afterNonZeroIdx := getNonZeroIdx(nums, firstZeroIdx+1)

	for firstZeroIdx != -1 && afterNonZeroIdx != -1 {
		nums[firstZeroIdx], nums[afterNonZeroIdx] = nums[afterNonZeroIdx], nums[firstZeroIdx]
		firstZeroIdx = getZeroIdx(nums, firstZeroIdx+1)
		afterNonZeroIdx = getNonZeroIdx(nums, afterNonZeroIdx+1)
	}
}

func getZeroIdx(nums []int, startIdx int) int {
	for i := startIdx; i < len(nums); i++ {
		if nums[i] == 0 {
			return i
		}
	}
	return -1
}

func getNonZeroIdx(nums []int, startIdx int) int {
	for i := startIdx; i < len(nums); i++ {
		if nums[i] != 0 {
			return i
		}
	}
	return -1
}
