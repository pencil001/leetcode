package main

import (
	"fmt"
)

func main() {
	fmt.Println(numIslands([][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}) == 1)

	fmt.Println(numIslands([][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}) == 3)
}

func numIslands(grid [][]byte) int {
	count := 0
	n := len(grid)
	if n == 0 {
		return count
	}
	m := len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				mark(grid, i, j, n, m)
				count++
			}
		}
	}
	return count
}

func mark(grid [][]byte, i, j, n, m int) {
	if i < 0 || j < 0 || i >= n || j >= m {
		return
	}
	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'

	mark(grid, i-1, j, n, m)
	mark(grid, i+1, j, n, m)
	mark(grid, i, j+1, n, m)
	mark(grid, i, j-1, n, m)
}
