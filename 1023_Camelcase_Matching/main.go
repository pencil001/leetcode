package main

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"))
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBaT"))
	fmt.Println(camelMatch([]string{"CompetitiveProgramming", "CounterPick", "ControlPanel"}, "CooP"))
}

func camelMatch(queries []string, pattern string) []bool {
	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = camelMatchOne(q, pattern)
	}
	return res
}

func camelMatchOne(query string, pattern string) bool {
	if len(pattern) > len(query) {
		return false
	}
	if len(pattern) == 0 {
		return true
	}
	if len(query) == 0 && len(pattern) == 0 {
		return true
	}

	i, j := 0, 0
	for i < len(query) && j < len(pattern) {
		if query[i] == pattern[j] {
			i++
			j++
			continue
		}

		// query[i] != pattern[j]
		if !isUpper(query[i]) {
			i++
			continue
		} else {
			return false
		}
	}

	if j < len(pattern) {
		return false
	}

	if i < len(query) {
		for ; i < len(query); i++ {
			if isUpper(query[i]) {
				return false
			}
		}
	}

	return true
}

func isUpper(c byte) bool {
	return c >= 65 && c <= 90
}
