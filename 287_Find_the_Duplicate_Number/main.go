package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 2, 3, 4}
	fmt.Println(findDuplicate3(nums))
}

// for just one duplicate
func findDuplicate(nums []int) int {
	n := len(nums) - 1

	xor := 0
	for i := 0; i < n; i++ {
		xor ^= i + 1
	}

	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}

	return xor
}

func findDuplicate2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}

func findDuplicate3(nums []int) int {
	slow := 0
	fast := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
