package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("234"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mapper := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	result := []string{""}
	for i := 0; i < len(digits); i++ {
		b := digits[i]
		values := mapper[b]

		tmp := []string{}
		for j := 0; j < len(result); j++ {
			for k := 0; k < len(values); k++ {
				s1 := result[j]
				s2 := string(values[k])
				tmp = append(tmp, s1+s2)
			}
		}
		result = tmp
	}
	return result
}
