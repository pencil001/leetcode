package p0200

func numIslandsDfs2(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	total := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == '1' {
				total++
				dfs2(grid, m, n, i, j)
			}
		}
	}

	return total
}

func dfs2(grid [][]byte, m, n int, i, j int) {
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	if i-1 >= 0 {
		dfs2(grid, m, n, i-1, j)
	}
	if i+1 < m {
		dfs2(grid, m, n, i+1, j)
	}
	if j-1 >= 0 {
		dfs2(grid, m, n, i, j-1)
	}
	if j+1 < n {
		dfs2(grid, m, n, i, j+1)
	}
}
