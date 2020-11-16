package p0003

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	window := map[uint8]int{}

	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right++

		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		res = max(res, right-left)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

//leetcode submit region end(Prohibit modification and deletion)
