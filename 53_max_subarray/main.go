package main

import "fmt"

func maxSubArray_divide(nums []int) int {
	ans := 0
	memoize := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rs := dfs(nums, i, memoize)
		ans = max(ans, rs)
	}

	return ans
}

func dfs(nums []int, i int, memoize map[int]int) int {
	if i == 0 {
		memoize[i] = nums[i]
		return nums[i]
	}

	if i == 1 {
		ans := max(nums[i], nums[i]+nums[0])
		memoize[i] = ans
		return ans
	}

	mem := memoize[i-1]
	ans := max(nums[i], nums[i]+mem)
	memoize[i] = ans

	return ans
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ans := nums[0]
	rs := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = max(nums[i], nums[i]+ans)
		if ans > rs {
			rs = ans
		}
	}

	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}
