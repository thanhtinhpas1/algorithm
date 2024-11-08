package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])

	return merge(first, second)
}

func merge(arr1, arr2 []int) []int {
	var rs []int
	var i, j int

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			rs = append(rs, arr1[i])
			i++
		} else {
			rs = append(rs, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		rs = append(rs, arr1[i])
		i++
	}

	for j < len(arr2) {
		rs = append(rs, arr2[j])
		j++
	}

	return rs
}

func main() {

	arr := []int{2, 1, 6, 4, 6, 9, 9, 4, 2, 7, 11, 44, 66, 44, 2346, 24754, 8653, 34, 8645, 2341, 22, 3, 4, 1, 6, 8}
	fmt.Println(mergeSort(arr))
}
