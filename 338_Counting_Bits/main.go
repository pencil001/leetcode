package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBits2(5))
}

func countBits(num int) []int {
	res := []int{}
	for i := 0; i <= num; i++ {
		n := i
		t := 0
		for n != 0 {
			n = n & (n - 1)
			t++
		}
		res = append(res, t)
	}
	return res
}

func countBits2(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	offset := 1
	for i := 1; i <= num; i++ {
		if offset*2 == i {
			offset *= 2
		}

		res[i] = res[i-offset] + 1
	}
	return res
}

func countBits3(num int) []int {
	array := make([]int, num+1)
	array[0] = 0
	for i := 1; i <= num; i++ {
		count := i & 0x1
		array[i] = array[i>>1] + count
	}
	return array
}
