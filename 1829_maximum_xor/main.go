package main

import (
	"fmt"
)

func getMaximumXor(nums []int, maximumBit int) []int {
	maxXor := (1 << maximumBit) - 1
	xor := nums[0]	

	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}

	rs := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		rs[i] = maxXor ^ xor
		xor ^= nums[len(nums)-i-1]
	}

	return rs
}

func main() {
	arr := []int{0, 1, 1, 3}
	maxBit := 2

	fmt.Println(getMaximumXor(arr, maxBit)) // [0,3,2,3]
}
