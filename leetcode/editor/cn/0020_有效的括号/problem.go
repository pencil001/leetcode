package p0020

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	match := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
			continue
		}

		if (len(stack) == 0) || c != match[stack[len(stack)-1]] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
