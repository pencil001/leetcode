package main

import "fmt"

func main() {
	strs := []string{
		"3[a]2[bc]",
		"3[a2[c]]",
		"2[abc]3[cd]ef",
	}
	for _, s := range strs {
		fmt.Println(decodeString(s))
	}
}

func decodeString(s string) string {
	res := []byte{}
	braces := []int{}
	repeats := []int{}

	repeatNum := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			repeatNum = 10*repeatNum + int(s[i]-'0')
		} else if s[i] == '[' {
			braces = append(braces, len(res))
			repeats = append(repeats, repeatNum)
			repeatNum = 0
		} else if s[i] == ']' {
			idxBrace := braces[len(braces)-1]
			rptNum := repeats[len(repeats)-1]
			remain := res[:idxBrace]
			rpt := res[idxBrace:]
			res = remain
			for j := 0; j < rptNum; j++ {
				res = append(res, rpt...)
			}
			braces = braces[:len(braces)-1]
			repeats = repeats[:len(repeats)-1]
		} else {
			res = append(res, s[i])
		}
	}

	return string(res[:])
}
