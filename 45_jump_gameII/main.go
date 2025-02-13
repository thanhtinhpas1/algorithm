package main

import (
	"fmt"
)

func jump(nums []int) int {
	near, far, jumps := 0, 0, 0

	for far < len(nums)-1 {
		farthest := 0
		for i := near; i <= far; i++ {
			farthest = max(farthest, i+nums[i])
		}

		near = far + 1
		far = farthest
		jumps++
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
