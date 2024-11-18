package main

import "fmt"

func findRepeatedDnaSequences_loop(s string) []string {
	var ans []string
	m := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		substr := s[i : i+10]
		frequency, exists := m[substr]

		if exists {
			if frequency == 1 {
				ans = append(ans, substr)
				m[substr] = m[substr] + 1
			}
		} else {
			m[substr] = 1
			ans = append(ans, substr)
		}
	}

	return ans
}

// sliding window
func findRepeatedDnaSequences(s string) []string {
	i := 0
	j := i + 10
	hash := make(map[string]int)
	var ans []string

	for i < j && j < len(s) {
		substr := s[i:j]
		frequency := hash[substr]

		if frequency == 0 {
			hash[substr] = 1
		} else {
			if frequency == 1 {
				ans = append(ans, substr)
			}
			hash[substr] = hash[substr] + 1
		}

		i++
		j = i + 10
	}

	return ans
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
