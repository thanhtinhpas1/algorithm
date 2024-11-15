package main

import "fmt"

func findLengthOfShortestSubarray(arr []int) int {
	right := len(arr) - 1
	for right > 0 {
		if arr[right-1] <= arr[right] {
			right--
		} else {
			break
		}
	}

	if right == 0 {
		return right
	}

	ans := right
	left := 0

	for left < right && (left == 0 || arr[left-1] <= arr[left]) {
		for right < len(arr) && arr[left] > arr[right] {
			right++
		}

		ans = min(ans, right-left-1)
		left++
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
	arr := []int{1, 2, 3, 3, 10, 1, 3, 3, 5}
	fmt.Println(findLengthOfShortestSubarray(arr))
}
