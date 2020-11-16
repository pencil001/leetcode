package p0076

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	window := map[uint8]int{}
	need := map[uint8]int{}

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	offset, length := 0, math.MaxInt32

	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++

		window[c]++
		if window[c] == need[c] {
			valid++
		}

		for valid == len(need) {
			l := right - left
			if l < length {
				length = l
				offset = left
			}

			d := s[left]
			left++

			if window[d] == need[d] {
				valid--
			}
			window[d]--
		}
	}

	if length != math.MaxInt32 {
		return s[offset : offset+length]
	}

	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
