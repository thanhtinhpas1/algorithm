package main

import "fmt"

const maxInt = int(^uint(0) >> 1)

func minSubArrayLen(target int, nums []int) int {
	ans := maxInt
	i := 0
	j := 1
	sum := nums[i]

	for i < len(nums) && j < len(nums) {
		sum += nums[j]

		for i <= j && sum >= target {
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}

		j++
	}

	if ans == maxInt {
		return 0
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, arr)) // 4,3
}
