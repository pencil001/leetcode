package main

import (
	"fmt"
)

func main() {
	s := "cbaebabacd"
	p := "abc"
	// s := "aa"
	// p := "bb"
	// s := "abab"
	// p := "ab"
	pos := findAnagrams(s, p)
	fmt.Println(pos)
}

func findAnagrams(s string, p string) []int {
	res := []int{}
	if len(s) < len(p) {
		return res
	}

	dict := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		dict[p[i]]++
	}
	counter := len(dict)

	begin, end := 0, 0
	for end < len(s) {
		e := s[end]
		if _, ok := dict[e]; ok {
			dict[e]--
			if dict[e] == 0 {
				counter--
			}
		}
		end++

		for counter == 0 {
			b := s[begin]
			if _, ok := dict[b]; ok {
				dict[b]++
				if dict[b] > 0 {
					counter++
				}
			}
			if end-begin == len(p) {
				res = append(res, begin)
			}
			begin++
		}
	}

	return res
}
