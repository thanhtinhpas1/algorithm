package main

import "fmt"

func longestSubstring(s string, k int) int {
	return calc(s, 0, len(s), k)
}

func calc(s string, start, end, k int) int {
	if end < k {
		return 0
	}

	idx := make(map[rune]int)

	for i := start; i < end; i++ {
		idx[rune(s[i]-'a')] += 1
	}

	for mid := start; mid < end; mid++ {
		if idx[rune(s[mid]-'a')] >= k {
			continue
		}

		midNext := mid + 1
		for midNext < end && idx[rune(s[midNext]-'a')] < k {
			midNext++
		}

		return max(calc(s, start, mid, k), calc(s, midNext, end, k))
	}

	return end - start
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(longestSubstring("aaabb", 3))  // output: 3
	fmt.Println(longestSubstring("ababbc", 2)) // output: 5
}
