package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)

	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		key := string(runes)
		if _, ok := hash[key]; ok {
			hash[key] = append(hash[key], s)
		} else {
			hash[key] = []string{s}
		}
	}

	res := make([][]string, 0)
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}
