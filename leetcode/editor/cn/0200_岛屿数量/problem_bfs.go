package p0200

func numIslandsBfs(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	total := 0

	queue := make([]Key, 0)
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] == '1' {
				total++
				grid[i][j] = '0'

				var k Key
				queue = append(queue, Key{i, j})
				for len(queue) > 0 {
					k, queue = queue[0], queue[1:]

					if k.X-1 >= 0 && grid[k.X-1][k.Y] == '1' {
						grid[k.X-1][k.Y] = '0'
						queue = append(queue, Key{k.X - 1, k.Y})
					}
					if k.X+1 < m && grid[k.X+1][k.Y] == '1' {
						grid[k.X+1][k.Y] = '0'
						queue = append(queue, Key{k.X + 1, k.Y})
					}
					if k.Y-1 >= 0 && grid[k.X][k.Y-1] == '1' {
						grid[k.X][k.Y-1] = '0'
						queue = append(queue, Key{k.X, k.Y - 1})
					}
					if k.Y+1 < n && grid[k.X][k.Y+1] == '1' {
						grid[k.X][k.Y+1] = '0'
						queue = append(queue, Key{k.X, k.Y + 1})
					}
				}
			}
		}
	}

	return total
}
