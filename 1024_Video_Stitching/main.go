package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(videoStitching([][]int{[]int{0, 2}, []int{4, 6}, []int{8, 10}, []int{1, 9}, []int{1, 5}, []int{5, 9}}, 10))
	fmt.Println(videoStitching([][]int{[]int{0, 1}, []int{1, 2}}, 5))
	fmt.Println(videoStitching([][]int{[]int{0, 1}, []int{6, 8}, []int{0, 2}, []int{5, 6}, []int{0, 4}, []int{0, 3}, []int{6, 7}, []int{1, 3}, []int{4, 7}, []int{1, 4}, []int{2, 5}, []int{2, 6}, []int{3, 4}, []int{4, 5}, []int{5, 7}, []int{6, 9}}, 9))
	fmt.Println(videoStitching([][]int{[]int{0, 4}, []int{2, 8}}, 5))
	fmt.Println(videoStitching([][]int{[]int{5, 7}, []int{1, 8}, []int{0, 0}, []int{2, 3}, []int{4, 5}, []int{0, 6}, []int{5, 10}, []int{7, 10}}, 5))
}

func videoStitching(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}
	if len(clips) == 0 {
		return -1
	}
	sort.Sort(ByStart(clips))

	if clips[0][0] != 0 {
		return -1
	}
	slims := [][]int{}
	curX := clips[0][0]
	curY := clips[0][1]
	for i := 0; i < len(clips); i++ {
		if clips[i][0] == curX {
			curY = clips[i][1]
			continue
		} else {
			slims = append(slims, []int{curX, curY})
			curX = clips[i][0]
			curY = clips[i][1]
		}
	}
	slims = append(slims, []int{curX, curY})

	fmt.Println(slims)

	end := slims[0][1]
	sum := 1
	for {
		if end >= T {
			break
		}

		toEnd := end
		for _, v := range slims {
			if v[0] <= end && v[1] > toEnd {
				toEnd = v[1]
			}
		}
		if toEnd > end {
			sum++
			end = toEnd
		} else {
			return -1
		}
	}

	return sum
}

type ByStart [][]int

func (s ByStart) Len() int      { return len(s) }
func (s ByStart) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByStart) Less(i, j int) bool {
	return (s[i][0] < s[j][0]) || ((s[i][0] == s[j][0]) && (s[i][1] < s[j][1]))
}
