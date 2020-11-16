package p0567

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	window := map[uint8]int{}
	need := map[uint8]int{}
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++

		window[c]++
		if window[c] == need[c] {
			valid++
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[left]
			left++

			if window[d] == need[d] {
				valid--
			}
			window[d]--
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
