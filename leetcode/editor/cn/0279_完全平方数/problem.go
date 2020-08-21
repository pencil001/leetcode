package p0279

//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	step := 0
	visited := make([]bool, n+1)
	visited[0] = true

	queue := []int{0}
	for len(queue) > 0 {
		sz := len(queue)
		step++

		for i := 0; i < sz; i++ {
			t := 0
			t, queue = queue[0], queue[1:]
			for j := 1; j*j <= n-t; j++ {
				s := t + j*j
				if s == n {
					return step
				}

				if !visited[s] {
					visited[s] = true
					queue = append(queue, s)
				}
			}
		}
	}

	return step
}

//leetcode submit region end(Prohibit modification and deletion)
