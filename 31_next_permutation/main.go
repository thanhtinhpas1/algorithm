package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	idx1 := -1
	idx2 := -1

	for 
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func reverse(nums []int, start int) {
	end := len(nums) - 1
	for start <= end {
		swap(nums, start, end)
		start++
		end--
	}
}

func main() {
	arr := []int{4, 2, 0, 2, 3, 2, 0}
	nextPermutation(arr)
	fmt.Println(arr) // 1, 3, 2
}
