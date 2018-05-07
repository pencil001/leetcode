package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	ans := []string{}
	backtrack(&ans, "", 0, 0, n)
	return ans
}

func backtrack(ans *[]string, str string, open int, close int, max int) {
	if len(str) == 2*max {
		*ans = append(*ans, str)
		return
	}

	if open < max {
		backtrack(ans, str+"(", open+1, close, max)
	}
	if close < open {
		backtrack(ans, str+")", open, close+1, max)
	}
}
