package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(str []string) [][]string {
	// anagrams rephase by origin characters to a new word

	result := [][]string{}
	hash := make(map[string][]string)

	for _, word := range str {
		chars := append([]rune(nil), []rune(word)...)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		hash[string(chars)] = append(hash[string(chars)], word)
	}

	for _, v := range hash {
		result = append(result, v)
	}

	return result
}

func main() {
	in := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("%v", groupAnagrams(in))

	// [["bat"],["nat","tan"],["ate","eat","tea"]]
}
