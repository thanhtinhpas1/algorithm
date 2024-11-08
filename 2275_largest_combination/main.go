package main

import "fmt"

func largestCombination(candidates []int) int {
	if len(candidates) == 1 {
		return 1
	}

	rs := make([]int, 24)
	max := 0

	for _, candidate := range candidates {
		i := 23
		for candidate > 0 && i >= 0 {
			if candidate&1 == 1 {
				rs[i]++
			}

			if rs[i] > max {
				max = rs[i]
			}

			candidate = candidate >> 1
			i--
		}
	}

	return max
}

func main() {
	// 16 & 17 & 62 & 24
	// 12: 		 1100
	// 16:    10000
	// 17:    10001
	// 62:   111110
	// 24:    11000
	// 71:  1000111
	arr := []int{16, 17, 71, 62, 12, 24, 14}
	fmt.Println(largestCombination(arr)) // 4
}
