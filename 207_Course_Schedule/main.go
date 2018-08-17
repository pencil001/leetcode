package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFinish(2, [][]int{[]int{1, 0}}) == true)
	fmt.Println(canFinish(2, [][]int{[]int{1, 0}, []int{0, 1}}) == false)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	degrees := make([]int, numCourses)
	for _, p := range prerequisites {
		degrees[p[0]]++
	}

	queue := []int{}
	for k, d := range degrees {
		if d == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) != 0 {
		v := queue[0]
		for _, p := range prerequisites {
			if p[1] == v {
				k := p[0]
				degrees[k]--
				if degrees[k] == 0 {
					queue = append(queue, k)
				}
			}
		}
		queue = queue[1:]
	}

	for _, d := range degrees {
		if d != 0 {
			return false
		}
	}
	return true
}
