package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	arr := []rune(s)
	i := 0
	j := 1
	ans := 1
	m := make(map[rune]int)

	for i < j && j < len(arr) {
		m[arr[i]] = i

		if _, exists := m[arr[j]]; !exists {
			m[arr[j]] = j

			if len(m) > ans {
				ans = len(m)
			}

			j++
			continue
		}

		i = m[arr[j]]+1
		j = i + 1
		m = make(map[rune]int)
	}

	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // 3
}
