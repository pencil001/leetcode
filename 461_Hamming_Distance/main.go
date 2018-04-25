package main

import "fmt"

func main() {
	x := 1577962638
	y := 1727613287
	z := hammingDistance(x, y)
	fmt.Print(z)
}

func hammingDistance(x int, y int) int {
	diff := 0
	for i := 0; i < 32; i++ {
		shift := 1 << uint32(i)
		if ((x & shift) ^ (y & shift)) != 0 {
			diff++
		}
	}
	return diff
}
