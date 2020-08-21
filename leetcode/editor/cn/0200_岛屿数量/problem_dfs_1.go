package p0200

//leetcode submit region begin(Prohibit modification and deletion)
type Key struct {
	X, Y int
}

func numIslandsDfs1(grid [][]byte) int {
	visited := make(map[Key]bool)
	m := len(grid)
	n := 0
	if m > 0 {
		n = len(grid[0])
	}
	total := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == '1' && !visited[Key{i, j}] {
				dfs1(grid, m, n, i, j, &visited)
				total++
			}
		}
	}

	return total
}

func dfs1(grid [][]byte, m, n int, i, j int, visited *map[Key]bool) {
	if grid[i][j] == '0' || (*visited)[Key{i, j}] {
		return
	}

	(*visited)[Key{i, j}] = true
	if i-1 >= 0 {
		dfs1(grid, m, n, i-1, j, visited)
	}
	if i+1 < m {
		dfs1(grid, m, n, i+1, j, visited)
	}
	if j-1 >= 0 {
		dfs1(grid, m, n, i, j-1, visited)
	}
	if j+1 < n {
		dfs1(grid, m, n, i, j+1, visited)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
