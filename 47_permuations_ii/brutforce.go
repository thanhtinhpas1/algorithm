package main

import (
	"fmt"
	"sort"
	"strings"
)

func permuteUnique_1(nums []int) [][]int {
	var result [][]int
	frequence := make(map[int]int)

	for _, num := range nums {
		frequence[num] = frequence[num] + 1
	}

	var backtrack func(nums []int, tmp []int)
	sort.Ints(nums)

	backtrack = func(nums []int, tmp []int) {
		if len(tmp) == len(nums) {
			result = append(result, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if isEnoughFrequence(tmp, nums[i], frequence) {
				continue
			}

			tmp = append(tmp, nums[i])

			backtrack(nums, append([]int(nil), tmp...))
			tmp = tmp[:len(tmp)-1]
		}
	}

	backtrack(nums, []int{})

	// filter result
	unique := make(map[string]bool)
	var uniqueResult [][]int

	for _, arr := range result {
		strBuilder := strings.Builder{}
		for _, num := range arr {
			strBuilder.WriteString(fmt.Sprintf("%v", num))
		}

		key := strBuilder.String()

		if unique[key] {
			continue
		}

		uniqueResult = append(uniqueResult, arr)
		unique[key] = true
	}

	return uniqueResult
}

func isEnoughFrequence(arr []int, num int, frequence map[int]int) bool {
	count := 0
	for _, item := range arr {
		if item == num {
			count++
		}
	}

	return count >= frequence[num]
}

// func main() {
// 	nums := []int{1, 1, 0, 0, 1, -1, -1, 1}
// 	/*
// 			[[1,1,2],
// 		 	[1,2,1],
// 		 	[2,1,1]]
// 	*/

// 	fmt.Printf("%v", permuteUnique(nums))
// }
