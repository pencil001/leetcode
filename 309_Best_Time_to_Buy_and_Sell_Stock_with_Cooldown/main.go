package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	s0 := make([]int, n)
	s1 := make([]int, n)
	s2 := make([]int, n)
	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = math.MinInt32
	for i := 1; i < n; i++ {
		s0[i] = max(s0[i-1], s2[i-1])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}

	return max(s0[n-1], s2[n-1])
}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := make([]int, n+1)
	buy[1] = 0 - prices[0]
	sel := make([]int, n+1)
	for i := 2; i <= n; i++ {
		buy[i] = max(buy[i-1], sel[i-2]-prices[i-1])
		sel[i] = max(sel[i-1], buy[i-1]+prices[i-1])
	}
	return sel[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
