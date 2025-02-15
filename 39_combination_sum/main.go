package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var backtrack func(start int, combinations []int, remain int)

	backtrack = func(start int, combinations []int, remain int) {
		if remain < 0 {
			return
		}

		if remain == 0 {
			result := make([]int, len(combinations))
			copy(result, combinations)
			ans = append(ans, result)
			return
		}

		for i := start; i < len(candidates); i++ {
			combinations = append(combinations, candidates[i])
			backtrack(i, combinations, remain-candidates[i])
			combinations = combinations[:len(combinations)-1]
		}
	}

	backtrack(0, []int{}, target)
	return ans
}

func main() {
	arr := []int{2, 3, 5}
	fmt.Printf("%v", combinationSum(arr, 8))
}
