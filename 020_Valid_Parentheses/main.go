package main

import (
	"fmt"
)

func main() {
	test := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"]]]]]",
	}
	for _, s := range test {
		ok := isValid(s)
		fmt.Println(ok)
	}
}

func isValid(s string) bool {
	index := -1
	stack := make([]rune, len(s))
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			index++
			stack[index] = c
		} else {
			if index < 0 {
				return false
			}
			v := stack[index]
			index--
			if (c == ')' && v != '(') || (c == ']' && v != '[') || (c == '}' && v != '{') {
				return false
			}
		}
	}
	return index == -1
}
