package main

import "fmt"

func letterCombinations(digits string) []string {
	var mapping = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	var ans []string
	var backtracking func(idx int, combination string)

	backtracking = func(idx int, combination string) {
		if idx == len(digits) {
			ans = append(ans, combination)
			return
		}

		letters := mapping[digits[idx]]
		for _, letter := range letters {
			backtracking(idx+1, combination+letter)
		}
	}

	backtracking(0, "")

	return ans
}

func main() {
	fmt.Println(fmt.Sprintf("%v", letterCombinations("23")))
}
