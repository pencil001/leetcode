package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12) == 3)
	fmt.Println(numSquares(13) == 2)
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		min := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j]+1 < min {
				min = dp[i-j*j] + 1
			}
		}
		dp[i] = min
	}
	return dp[n]
}
