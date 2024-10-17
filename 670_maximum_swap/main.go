package main

import (
	"fmt"
	"sort"
)

func maximumSwap(origin int) int {
	var arr []int
	var biggest, biggestPosition int

	num := origin
	// 1234

	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}

	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)

	sort.Ints(sortedArr)

	i, j := len(arr)-1, len(arr)-1
	for arr[i] == sortedArr[j] && i > 0 && j > 0 {
		i--
		j--
	}

	swapArr := arr[:j+1]

	for i := 0; i < len(swapArr); i++ {
		if swapArr[i] > biggest {
			biggest = swapArr[i]
			biggestPosition = i
		}
	}

	// 4 3 2 1
	for i := len(swapArr) - 1; i >= 0; i-- {
		if i == biggestPosition {
			break
		}

		if arr[i] < biggest {
			arr[i], arr[biggestPosition] = arr[biggestPosition], arr[i]
			break
		}
	}

	var result int
	for i := len(arr) - 1; i >= 0; i-- {
		result = result*10 + arr[i]
	}

	return result
}

func main() {
	fmt.Println(maximumSwap(9973))
}
