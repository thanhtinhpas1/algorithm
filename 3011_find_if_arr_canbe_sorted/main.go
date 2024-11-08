package main

import "fmt"

type ArrMinMax struct {
	Max int
	Min int
}

func canSortArray(nums []int) bool {
	if isSorted(nums) {
		return true
	}

	idx := make([]int, len(nums))

	for i, num := range nums {
		idx[i] = calcSetBit(num)
	}

	minMaxArr := []ArrMinMax{}
	curMin := nums[0]
	curMax := nums[0]
	curBit := idx[0]

	for i, num := range nums {
		if idx[i] != curBit {
			minMaxArr = append(minMaxArr, ArrMinMax{Min: curMin, Max: curMax})
			curBit = idx[i]
			curMax = num
			curMin = num

			continue
		}

		if num > curMax {
			curMax = num
		}

		if num < curMin {
			curMin = num
		}
	}
	minMaxArr = append(minMaxArr, ArrMinMax{Min: curMin, Max: curMax})

	for i := 0; i < len(minMaxArr)-1; i++ {
		if minMaxArr[i].Max > minMaxArr[i+1].Min {
			return false
		}
	}

	return true
}

func calcSetBit(num int) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}

		num = num >> 1
	}

	return count
}

func isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{2, 4, 8}
	fmt.Println(canSortArray(arr))

	arr = []int{8, 4, 2, 30, 15}
	fmt.Println(canSortArray(arr))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(canSortArray(arr))

	arr = []int{3, 16, 8, 4, 2}
	fmt.Println(canSortArray(arr))

	arr = []int{21, 17}
	fmt.Println(canSortArray(arr))

	arr = []int{174, 175, 234, 188}
	fmt.Println(canSortArray(arr))
}
