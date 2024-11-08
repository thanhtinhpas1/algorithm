package main

import (
	"fmt"
	"sort"
)

// func threeSum(nums []int) [][]int {
// 	idx := make(map[int]int)

// 	for _, num := range nums {
// 		_, exists := idx[num]
// 		if !exists {
// 			idx[num] = 1
// 		} else {
// 			idx[num]++
// 		}
// 	}

// 	var rs [][]int
// 	dupPrevent := make(map[string]bool)

// 	i := 0
// 	j := i + 1
// 	for i < len(nums)-2 {
// 		if j == len(nums)-1 {
// 			i++
// 			j = i + 1
// 		}

// 		sum := nums[i] + nums[j]

// 		expect := 0 - sum
// 		fr := idx[expect]

// 		dupNums := 0
// 		if nums[i] == expect {
// 			dupNums++
// 		}

// 		if nums[j] == expect {
// 			dupNums++
// 		}

// 		if fr > dupNums {
// 			expectArr := []int{nums[i], nums[j], expect}
// 			sort.Ints(expectArr)
// 			key := fmt.Sprintf("%d:%d:%d", expectArr[0], expectArr[1], expectArr[2])
// 			if !dupPrevent[key] {
// 				rs = append(rs, expectArr)
// 				dupPrevent[key] = true
// 			}
// 		}

// 		j++
// 	}

// 	return rs
// }

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var rs [][]int

	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				rs = append(rs, []int{nums[i], nums[j], nums[k]})

				for j < k && nums[j] == nums[j+1] {
					j++
				}

				for j < k && nums[k] == nums[j-1] {
					k--
				}

				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}


	return rs
}

func main() {
	// -4, -1, -1, 0, 1, 2
	arr := []int{0, 0, 0, 0} // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum(arr))
}
