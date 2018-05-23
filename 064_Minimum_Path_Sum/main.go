package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := 0
			if i-1 < 0 && j-1 < 0 {
				v = 0
			} else if i-1 < 0 {
				v = dp[i][j-1]
			} else if j-1 < 0 {
				v = dp[i-1][j]
			} else {
				v = min(dp[i-1][j], dp[i][j-1])
			}
			dp[i][j] = v + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minPathSum2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}
