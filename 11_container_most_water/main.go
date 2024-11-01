package main

import (
	"fmt"
)

func maxArea(arr []int) int {
	var max int

	i := 0
	j := len(arr) - 1
	for i < j {
		rs := (j - i) * min(arr[i], arr[j])
		if rs > max {
			max = rs
		}

		if arr[i] < arr[j] {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	rs := maxArea(arr)
	fmt.Println(rs)
}
