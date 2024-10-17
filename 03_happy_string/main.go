package main

import (
	"fmt"
	"sort"
)

type Item struct {
	key  int
	char rune
}

func longestDiverseString(a int, b int, c int) string {
	var cols = []Item{{a, rune('a')}, {b, rune('b')}, {c, rune('c')}}
	sort.SliceStable(cols, func(i, j int) bool {
		return cols[i].key > cols[j].key
	})

	a, b, c = cols[0].key, cols[1].key, cols[2].key

	var arr []rune
	var aBlock, bBlock, cBlock bool

	for (a > 0 && !aBlock) || (b > 0 && !bBlock) || (c > 0 && !cBlock) {
		if a > 0 && !aBlock {
			if len(arr) > 0 && arr[len(arr)-1] == cols[0].char {
				aBlock = true
			}

			if a > 2 {
				arr = append(arr, cols[0].char, cols[0].char)
				aBlock = true
				a -= 2
			} else {
				arr = append(arr, cols[0].char)
				a--
			}

			bBlock = false
			cBlock = false
		} else if b > 0 && !bBlock {
			if len(arr) > 0 && arr[len(arr)-1] == cols[1].char {
				bBlock = true
			}

			if b > 2 {
				arr = append(arr, cols[1].char, cols[1].char)
				b -= 2
				bBlock = true
			} else {
				arr = append(arr, cols[1].char)
				b--
			}

			aBlock = false
			cBlock = false
		} else if c > 0 && !cBlock {
			if len(arr) > 0 && arr[len(arr)-1] == cols[2].char {
				cBlock = true
			}

			if c > 2 {
				arr = append(arr, cols[2].char, cols[2].char)
				c -= 2
				cBlock = true
			} else {
				arr = append(arr, cols[2].char)
				c--
			}
			aBlock = false
			bBlock = false
		}
	}

	return string(arr)
}

func main() {
	fmt.Println(longestDiverseString(7, 1, 0))
	fmt.Println(longestDiverseString(1, 1, 7))
}
