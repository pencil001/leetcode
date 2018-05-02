package main

import (
	"fmt"
	"sort"
)

func main() {
	// [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
	people := [][]int{
		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1},
		[]int{5, 0},
		[]int{6, 1},
		[]int{5, 2},
	}
	fmt.Println(reconstructQueue(people))
}

func reconstructQueue(people [][]int) [][]int {
	res := [][]int{}

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	for _, p := range people {
		idx := p[1]
		res = append(res[:idx], append([][]int{p}, res[idx:]...)...)
	}

	return res
}

func reconstructQueue2(people [][]int) [][]int {
	res := [][]int{}

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	insert := func(idx int, person []int) {
		res = append(res, person)
		// 插入到末尾
		if len(res)-1 == idx {
			return
		}
		// 插入到中间位置
		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}

	for _, p := range people {
		insert(p[1], p)
	}

	return res
}
