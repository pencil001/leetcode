package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := findDisappearedNumbers2(nums)
	fmt.Println(res)
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			j := nums[i]
			if j == 0 {
				break
			}
			if nums[j-1] != j {
				nums[j-1], nums[i] = nums[i], nums[j-1]
			} else {
				break
			}
		}

		if nums[i] != i+1 {
			nums[i] = 0
		}
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func findDisappearedNumbers2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		j := abs(nums[i]) - 1
		if nums[j] > 0 {
			nums[j] = -nums[j]
		}
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}
