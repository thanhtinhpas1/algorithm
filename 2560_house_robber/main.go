package main

import (
	"fmt"
)

func minCapability(nums []int, k int) int {
	left, right := min(nums), max(nums)

	check := func(capability int) bool {
		count, i := 0, 0

		for i < len(nums) {
			if nums[i] <= capability {
				count++
				i += 2
			} else {
				i++
			}

			if count >= k {
				return true
			}
		}

		return false
	}

	for left < right {
		mid := (left + right) / 2

		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func min(nums []int) int {
	minNum := int(1e9)
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}

	return minNum
}

func max(nums []int) int {
	maxNum := 0

	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	return maxNum
}


func main() {
	houses := []int{2, 3, 5, 9}
	fmt.Println(minCapability(houses, 2)) // output: 5

	houses = []int{2, 7, 9, 3, 1}
	fmt.Println(minCapability(houses, 2)) // output: 2

	houses = []int{24, 109, 117, 142, 98, 94, 91, 130, 73, 48, 107, 77}
	fmt.Println(minCapability(houses, 5)) // output: 98
}
