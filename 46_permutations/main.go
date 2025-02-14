package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int

	var backtrack func(nums []int, tmp []int)

	backtrack = func(nums, tmp []int) {
		if len(tmp) == len(nums) {
			result = append(result, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if isContains(tmp, nums[i]) {
				continue
			}

			tmp = append(tmp, nums[i])
			newarr := []int{}
			copy(newarr, tmp)

			backtrack(nums, newarr)

			tmp = tmp[:len(tmp)-1]
		}
	}

	backtrack(nums, []int{})

	return result
}

func isContains(arr []int, x int) bool {
	for _, num := range arr {
		if num == x {
			return true
		}
	}

	return false
}

func main() {
	fmt.Printf("%v", permute([]int{1, 2, 3}))
}
