package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[l] == target {
			return l
		}

		if nums[r] == target {
			return r
		}

		if target > nums[l] && target < nums[mid] ||
			nums[l] > nums[mid] && target > nums[l] ||
			nums[l] > nums[mid] && target < nums[mid] {

			r = mid - 1
			continue
		}

		l = mid + 1
		continue
	}

	return -1
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	fmt.Println(search(arr, target))
}
