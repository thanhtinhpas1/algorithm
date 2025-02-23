package main

import (
	"fmt"
)

/**
iteration 1:
k = 0, i= 0, j =0
i = 0, j = 0 -> 3, k = 0
i = 0 -> 3, j = 3, k= 0
i = 3, j = 2 -> 0, k = 0
i = 2 -> 0, j = 0, k = 0

iteration 2:
k = 1
i =
**/

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}

	var i, j int
	N := len(matrix[0]) * len(matrix)
	k := 0

	for len(result) < N && i < len(matrix) && j < len(matrix[0]) {
		for ; len(result) < N && j < len(matrix[0])-k; j++ {
			result = append(result, matrix[i][j])
		}
		j--
		i++

		for ; len(result) < N && i < len(matrix)-k; i++ {
			result = append(result, matrix[i][j])
		}
		i--
		j--

		for ; len(result) < N && j >= k; j-- {
			result = append(result, matrix[i][j])
		}
		j++
		i--

		for ; len(result) < N && i > k; i-- {
			result = append(result, matrix[i][j])
		}

		i++
		j++
		k++
	}

	return result
}

func main() {
	matrix1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Printf("%v", spiralOrder(matrix1))
	// expect:
	// [1, 2, 3, 6, 9, 8, 7, 4, 5]
}
