package main

import "fmt"

// rotate matrix 90 degrees and not allocation new array
func rotate(matrix [][]int) {
	// matrix n
	n := len(matrix)

	for i := 0; i <= n/2; i++ {
		for j := i; j < n-1-i; j++ {
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}

func main() {
	arr := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(arr)
	fmt.Printf("%v", arr)
	/**
	7, 4, 1
	8, 5, 2
	9, 6, 3
	**/
}
