package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	var dp [100][100]int

	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < m; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}
