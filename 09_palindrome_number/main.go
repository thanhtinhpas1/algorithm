package main

import "fmt"

func isPalindrome(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}

	if x < 0 {
		return false
	}

	reverseNum := reverse(x)
	return reverseNum == x
}

func reverse(x int) int {
	var arr []int

	for x > 0 {
		arr = append(arr, x%10)

		x = x / 10
	}

	result := arr[0]

	for i := 1; i < len(arr); i++ {
		result = result * 10
		result += arr[i]
	}

	return result
}

func main() {
	fmt.Println(isPalindrome(121))
}
