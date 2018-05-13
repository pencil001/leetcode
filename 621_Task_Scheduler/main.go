package main

import (
	"fmt"
)

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}

func leastInterval(tasks []byte, n int) int {
	sum := len(tasks)
	if n == 0 {
		return sum
	}

	hash := make(map[byte]int)
	for i := 0; i < sum; i++ {
		hash[tasks[i]]++
	}

	maxV := 0
	maxT := 0
	for _, v := range hash {
		if v > maxV {
			maxV = v
			maxT = 0
		}
		if v == maxV {
			maxT++
		}
	}

	sum2 := (n+1)*(maxV-1) + maxT

	if sum > sum2 {
		return sum
	} else {
		return sum2
	}
}
