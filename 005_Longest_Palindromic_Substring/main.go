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
		fmt.Println(longestPalindrome3(s))
	}
}

// DP
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	lenS := len(s)
	dp := make([][]bool, lenS)
	start := 0
	end := 0
	for i := 0; i < lenS; i++ {
		dp[i] = make([]bool, lenS)
		dp[i][i] = true
		if i-1 >= 0 && s[i] == s[i-1] {
			dp[i-1][i] = true
			if i-(i-1) > end-start {
				end = i
				start = i - 1
			}
		}
	}

	for j := 0; j < lenS; j++ {
		for i := 0; i <= j; i++ {
			if dp[i][j] {
				if i-1 >= 0 && j+1 < lenS && s[i-1] == s[j+1] {
					dp[i-1][j+1] = true
					if j+1-(i-1) > end-start {
						end = j + 1
						start = i - 1
					}
				}
			}
		}
	}

	return s[start : end+1]
}

// DP
func longestPalindrome2(s string) string {
	start := 0
	end := 0
	dp := [1000][1000]bool{}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if i+1 < len(s) && s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			end = i + 1
		}
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if dp[i][j] {
				if i-1 >= 0 && j+1 < len(s) && s[i-1] == s[j+1] {
					dp[i-1][j+1] = true
					if j+1-(i-1) > end-start {
						end = j + 1
						start = i - 1
					}
				}
			}
		}
	}

	return s[start : end+1]
}

// Expand Around Center
func longestPalindrome3(s string) string {
	start := 0
	end := 0
	for c := 0; c <= 2*len(s)-1; c++ {
		left := c / 2
		right := left + c%2
		for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
			if right-left > end-start {
				start = left
				end = right
			}
			left--
			right++
		}
	}
	return s[start : end+1]
}
