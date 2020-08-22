package p0739

//leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	idx := []int{}

	for i, t := range T {
		for len(idx) > 0 && t > T[idx[len(idx)-1]] {
			last := idx[len(idx)-1]
			res[last] = i - last

			idx = idx[:len(idx)-1]
		}

		idx = append(idx, i)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
