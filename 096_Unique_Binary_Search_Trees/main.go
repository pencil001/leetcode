package main

import "fmt"

func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(5))
	fmt.Println(numTrees(19))
}

func numTrees(n int) int {
	hash := make(map[int]int)
	hash[0] = 1
	hash[1] = 1
	return numTreesEx(n, &hash)
}

func numTreesEx(n int, hash *map[int]int) int {
	if v, ok := (*hash)[n]; ok {
		return v
	}

	res := 0
	for i := 0; i < n; i++ {
		res += numTreesEx(i, hash) * numTreesEx(n-1-i, hash)
	}
	(*hash)[n] = res
	return res
}

func numTrees2(n int) int {
	hash := make(map[int]int)
	hash[0] = 1
	hash[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			hash[i] += hash[j] * hash[i-1-j]
		}
	}
	return hash[n]
}
