package main

import (
	"fmt"
)

func compressedString(s string) string {
	var result []rune
	word := []rune(s)
	count := 0
	curr := word[0]

	for _, ch := range word {
		if ch != curr {
			result = append(result, rune(count+'0'), curr)
			curr = ch
			count = 1

			continue
		}

		if count >= 9 {
			result = append(result, rune('9'), curr)
			count = 0
		}

		count++
		curr = ch
	}

	if count > 0 {
		result = append(result, rune(count+'0'), curr)
	}

	return string(result)
}

func main() {
	fmt.Println(compressedString("aaaaaaaaaaaaaabb")) // 9a5a2b
	fmt.Println(compressedString("mrm"))              // 1m1r1m
	fmt.Println(compressedString("aaaaaaaaay"))
}
