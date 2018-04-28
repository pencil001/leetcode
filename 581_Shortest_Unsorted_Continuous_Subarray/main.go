package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	// nums := []int{1, 2, 3, 4}
	fmt.Println(findUnsortedSubarray2(nums))
}

func findUnsortedSubarray3(nums []int) int {
	//Forward sweep
	foundPeak := false
	min := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if foundPeak {
			if nums[i] < min {
				min = nums[i]
			}
		} else {
			if (i+1) < len(nums) && nums[i+1] < nums[i] {
				foundPeak = true
			}
		}
	}

	if !foundPeak {
		return 0
	}

	//Backward Sweep
	foundValley := false
	max := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		if foundValley {
			if nums[i] > max {
				max = nums[i]
			}
		} else {
			if (i-1) >= 0 && nums[i-1] > nums[i] {
				foundValley = true
			}
		}
	}

	var minIndex, maxIndex int
	for i, num := range nums {
		if num > min {
			minIndex = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			maxIndex = i
			break
		}
	}
	return maxIndex - minIndex + 1
}

func findUnsortedSubarray2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	dup := make([]int, len(nums))
	copy(dup, nums)
	sort.Ints(dup)
	begin := len(nums)
	end := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != dup[i] {
			if i < begin {
				begin = i
			}
			if i > end {
				end = i
			}
		}
	}

	if end > begin {
		return end - begin + 1
	} else {
		return 0
	}
}

func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	dup := make([]int, len(nums))
	copy(dup, nums)
	sort.Ints(nums)
	begin := 0
	end := len(nums) - 1
	for begin < len(nums) {
		if nums[begin] == dup[begin] {
			begin++
		} else {
			break
		}
	}

	if begin == len(nums) {
		return 0
	}

	for {
		if nums[end] == dup[end] {
			end--
		} else {
			break
		}
	}
	return end - begin + 1
}
