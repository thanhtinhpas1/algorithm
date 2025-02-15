package main

import "fmt"

/**
// Complexity:
Chọn ra R phần tử từ N phần tử độc lập:
P(n, r) = n(n-1)(n02)...(n-r+1) = n! / (n-r)!
**/
func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	frequencyMap := make(map[int]int)

	for _, num := range nums {
		frequencyMap[num]++
	}

	var backtrack func(comb []int, counter int)
	backtrack = func(comb []int, counter int) {
		if len(comb) == len(nums) {
			result = append(result, append([]int(nil), comb...))
			return
		}

		for num, count := range frequencyMap {
			if count == 0 {
				continue
			}

			comb = append(comb, num)
			frequencyMap[num]--

			backtrack(comb, counter)

			frequencyMap[num]++
			comb = comb[:len(comb)-1]
		}
	}

	backtrack([]int{}, len(nums))
	return result
}

func main() {
	nums := []int{1, 1, 0, 0, 1, -1, -1, 1}
	/*
			[[1,1,2],
		 	[1,2,1],
		 	[2,1,1]]
	*/

	fmt.Printf("%v", permuteUnique(nums))
}
