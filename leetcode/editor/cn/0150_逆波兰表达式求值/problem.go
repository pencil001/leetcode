package p0150

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" {
			x := stack[len(stack)-2]
			y := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			res := 0
			switch t {
			case "+":
				res = x + y
			case "-":
				res = x - y
			case "*":
				res = x * y
			case "/":
				res = x / y
			}
			stack = append(stack, res)
		} else {
			n, _ := strconv.Atoi(t)
			stack = append(stack, n)
		}
	}

	return stack[len(stack)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
