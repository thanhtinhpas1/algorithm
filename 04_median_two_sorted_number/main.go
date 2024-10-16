package main

import (
	"fmt"
	"sort"
)

// That's brute force way, combine 2 arrays then sort and calculate the median
// Complexity: O((n + m) * log(n + m))
// Space: O(n+m)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)
	// O N(logN)
	sort.Ints(arr)

	// even
	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2-1]+arr[len(arr)/2]) / float64(2)
	}

	// odd
	return float64(arr[len(arr)/2])
}

// Using 2 pointers and move forward to copy sorted items in 2 arrays to new array, then find median
// Complexity: O(n+m)
// Space: O(n+m)
func findMedianSortedArrays_2(arr1, arr2 []int) float64 {
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}

	var i, j int
	var arr []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		arr = append(arr, arr1[i])
		i++
	}

	for j < len(arr2) {
		arr = append(arr, arr2[j])
		j++
	}

	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2-1]+arr[len(arr)/2]) / float64(2)
	}

	// odd
	return float64(arr[len(arr)/2])
}

// Using 2 pointers and move forward to copy sorted items in 2 arrays to new array, then find median
// Complexity: O(n+m)
// Space: O(1)
func findMedianSortedArrays_2_optimize(arr1, arr2 []int) float64 {
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}

	var i, j, m1, m2 int
	n := len(arr1)
	m := len(arr2)

	for count := 0; count <= (n+m)/2; count++ {
		m2 = m1
		if i != n && j != m {
			if arr1[i] > arr2[j] {
				m1 = arr2[j]
				j++
			} else {
				m1 = arr1[i]
				i++
			}
		} else if i < n {
			m1 = arr1[i]
			i++
		} else {
			m1 = arr2[j]
			j++
		}
	}

	if (n+m)%2 == 0 {
		return float64(m1+m2) / float64(2)
	}

	return float64(m1)
}

func main() {
	arr1 := []int{1, 3}
	arr2 := []int{2}

	arr3 := []int{1, 4}
	arr4 := []int{2, 3}

	fmt.Println(findMedianSortedArrays(arr1, arr2))
	fmt.Println(findMedianSortedArrays(arr3, arr4))

	fmt.Println(findMedianSortedArrays_2(arr1, arr2))
	fmt.Println(findMedianSortedArrays_2(arr3, arr4))

	fmt.Println(findMedianSortedArrays_2_optimize(arr1, arr2))
	fmt.Println(findMedianSortedArrays_2_optimize(arr3, arr4))
}
