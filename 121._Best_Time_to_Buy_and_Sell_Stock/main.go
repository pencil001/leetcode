package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	i := maxProfit(prices)
	fmt.Println(i)
}

func maxProfit(prices []int) int {
	begin, end, max := 0, 0, 0
	for end < len(prices) {
		v := prices[end] - prices[begin]
		if v <= 0 {
			begin = end
		}
		if v > max {
			max = v
		}
		end++
	}
	return max
}
