package main

import (
	"fmt"
)

func main() {
	n := climbStairs2(10)
	fmt.Println(n)
}

func climbStairs(n int) int {
	steps := []int{0}
	for i := 1; i <= n; i++ {
		step := 0
		if i-1 >= 0 {
			step += steps[i-1]
			if i-1 == 0 {
				step++
			}
		}
		if i-2 >= 0 {
			step += steps[i-2]
			if i-2 == 0 {
				step++
			}
		}
		steps = append(steps, step)
	}
	return steps[n]
}

func climbStairs2(n int) int {
	steps := make([]int, n+1)
	steps[0], steps[1], steps[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		steps[i] = steps[i-1] + steps[i-2]
	}
	return steps[n]
}
