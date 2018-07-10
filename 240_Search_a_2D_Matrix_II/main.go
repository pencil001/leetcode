package main

import (
	"fmt"
)

func main() {
	// matrix := [][]int{
	// 	[]int{1, 4, 7, 11, 15},
	// 	[]int{2, 5, 8, 12, 19},
	// 	[]int{3, 6, 9, 16, 22},
	// 	[]int{10, 13, 14, 17, 24},
	// 	[]int{18, 21, 23, 26, 30},
	// }
	matrix := [][]int{}
	ok := searchMatrix(matrix, 0)
	fmt.Println(ok)
}

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	row := 0
	col := m - 1
	for row < n && col >= 0 {
		v := matrix[row][col]
		if v == target {
			return true
		}
		if v > target {
			col--
		} else {
			row++
		}
	}

	return false
}
