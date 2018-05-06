package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 1, 1, 2, 2, 3}
	// k := 2
	nums := []int{-1, -1}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}

	freq := make([][]int, len(nums)+1)
	for k, v := range hash {
		l := freq[v]
		if l == nil {
			l = []int{}
		}
		freq[v] = append(freq[v], k)
	}

	res := []int{}
	for i := len(freq) - 1; i >= 0 && len(res) < k; i-- {
		l := freq[i]
		if l != nil {
			for j := 0; j < len(l) && len(res) < k; j++ {
				res = append(res, l[j])
			}
		}
	}
	return res
}
