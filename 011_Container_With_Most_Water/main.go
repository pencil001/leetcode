package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	fmt.Println(maxArea2([]int{}) == 0)
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	max := 0
	for i := 0; i < len(height); i++ {
		h1 := height[i]
		for j := i + 1; j < len(height); j++ {
			h2 := height[j]
			minH := h1
			if h2 < h1 {
				minH = h2
			}
			if (j-i)*minH > max {
				max = (j - i) * minH
			}
		}
	}
	return max
}

func maxArea2(height []int) int {
	maxA := 0
	l := 0
	r := len(height) - 1

	for l < r {
		h1 := height[l]
		h2 := height[r]
		h := min(h1, h2)
		maxA = max(maxA, h*(r-l))
		if h1 < h2 {
			l++
		} else {
			r--
		}
	}
	return maxA
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}

func max(v1, v2 int) int {
	if v1 < v2 {
		return v2
	} else {
		return v1
	}
}
