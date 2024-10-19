package main

import (
	"fmt"
)

func findKthBit(n int, k int) byte {
	return findN(n)[k-1]
}

func findN(n int) []byte {
	if n == 1 {
		return []byte{0}
	}

	if n == 2 {
		return []byte{0, 1, 1}
	}

	rs := findN(n - 1)
	rs = append(rs, byte(1))
	for i := len(rs) - 2; i >= 0; i-- {
		if rs[i] == 1 {
			rs = append(rs, 0)
		} else {
			rs = append(rs, 1)
		}
	}

	return rs
}

func main() {
	fmt.Println(findKthBit(4, 11))
}
