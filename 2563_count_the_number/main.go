package main

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	var rs int64
	for i, num := range nums {
		l := lowerBound(nums, i+1, len(nums), lower-num)
		u := upperBound(nums, i+1, len(nums), upper-num)

		rs += int64(u - l)
	}

	return rs
}

func lowerBound(nums []int, start, end, target int) int {
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func upperBound(nums []int, start, end, target int) int {
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func main() {
	arr := []int{0, 1, 7, 4, 4, 5}
	fmt.Println(countFairPairs(arr, 3, 6)) // 6
}
