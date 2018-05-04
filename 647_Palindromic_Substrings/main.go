package main

import (
	"fmt"
)

func main() {
	ss := []string{
		"babad",
		"cbbd",
		"abcba",
	}
	for _, s := range ss {
		fmt.Println(countSubstrings2(s))
	}
}

func countSubstrings(s string) int {
	res := 0
	for c := 0; c <= 2*len(s)-1; c++ {
		left := c / 2
		right := left + c%2
		for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
			res++
			left--
			right++
		}
	}
	return res
}

func countSubstrings2(s string) int {
	res := 0
	dp := [1000][1000]bool{}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		res++
		if i+1 < len(s) && s[i] == s[i+1] {
			dp[i][i+1] = true
			res++
		}
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if dp[i][j] {
				if i-1 >= 0 && j+1 < len(s) && s[i-1] == s[j+1] {
					if !dp[i-1][j+1] {
						dp[i-1][j+1] = true
						res++
					}
				}
			}
		}
	}

	return res
}
