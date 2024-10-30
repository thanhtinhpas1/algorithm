package main

import "fmt"

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	LIS := make([]int, n)
	LDS := make([]int, n)

	for i := range LIS {
		LIS[i] = 1
		LDS[i] = 1
	}

	// Compute LIS
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				LIS[i] = max(LIS[i], LIS[j]+1)
			}
		}
	}

	// Compute LDS
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] > nums[j] {
				LDS[i] = max(LDS[i], LDS[j]+1)
			}
		}
	}

	maxMountainLength := 0
	for i := 1; i < n-1; i++ {
		if LIS[i] > 1 && LDS[i] > 1 {
			maxMountainLength = max(maxMountainLength, LIS[i]+LDS[i]-1)
		}
	}

	return n - maxMountainLength
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	arr := []int{4, 5, 13, 17, 1, 7, 6, 11, 2, 8, 10, 15, 3, 9, 12, 14, 16}
	fmt.Println(minimumMountainRemovals(arr))
}

