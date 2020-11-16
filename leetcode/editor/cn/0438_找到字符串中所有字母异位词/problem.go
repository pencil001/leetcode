package p0438

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	window := map[uint8]int{}
	need := map[uint8]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	left, right := 0, 0
	valid := 0
	res := []int{}
	for right < len(s) {
		c := s[right]
		right++

		window[c]++
		if window[c] == need[c] {
			valid++
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}

			d := s[left]
			left++

			if window[d] == need[d] {
				valid--
			}
			window[d]--
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
