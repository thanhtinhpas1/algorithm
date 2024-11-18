package main

import "fmt"

func decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}

	var ans = make([]int, len(code))
	for i := 0; i < len(code); i++ {
		sum := 0

		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%len(code)]
			}
		}

		if k < 0 {
			for j := -1; j >= k; j-- {
				idx := (i + j + len(code)) % len(code)
 				sum += code[idx]
			}
		}

		ans[i] = sum
	}

	return ans
}

func main() {
	arr := []int{2, 4, 9, 3}
	fmt.Println(decrypt(arr, -2)) // output: 12,5,6,13
}
