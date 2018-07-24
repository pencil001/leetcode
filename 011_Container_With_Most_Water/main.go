package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
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
