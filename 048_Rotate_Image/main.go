package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	// matrix := [][]int{
	// 	[]int{5, 1, 9, 11},
	// 	[]int{2, 4, 8, 10},
	// 	[]int{13, 3, 6, 7},
	// 	[]int{15, 14, 12, 16},
	// }
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}
