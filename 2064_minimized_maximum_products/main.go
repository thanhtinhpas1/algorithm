package main

import (
	"fmt"
)

func minimizedMaximum(n int, quantities []int) int {
	max := 0
	for _, q := range quantities {
		if q > max {
			max = q
		}
	}

	left := 0
	right := max
	for left < right {
		mid := (left + right) / 2
		if canDistribute(n, quantities, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canDistribute(n int, quantities []int, k int) bool {
	j := 0
	remain := quantities[0]

	for i := 0; i < n; i++ {
		if remain <= k {
			j++

			if j == len(quantities) {
				return true
			} else {
				remain = quantities[j]
			}
		} else {
			remain = remain - k
		}
	}

	return false
}

func main() {
	fmt.Println(minimizedMaximum(6, []int{11, 6}))
}
