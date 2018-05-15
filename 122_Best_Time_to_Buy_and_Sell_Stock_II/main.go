package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	begin, end, curMax, totalMax := 0, 0, 0, 0
	for end < len(prices) {
		v := prices[end] - prices[begin]
		if v > curMax {
			totalMax = totalMax - curMax + v
			curMax = v
		} else {
			curMax = 0
			begin = end
		}
		end++
	}
	return totalMax
}
