package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	// unique quadruples: nums[a], nums[b], nums[c], nums[d]

	var ans [][]int
	sort.Ints(nums) // O(NLogN)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			expect := target - (nums[i] + nums[j])

			k := j + 1
			f := len(nums) - 1

			for k < f {
				if nums[k]+nums[f] == expect {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[f]})

					for k < f && nums[k] == nums[k+1] {
						k++
					}

					for f > k && nums[f] == nums[f-1] {
						f--
					}

					k++
					f--
					continue
				}

				if nums[k] < expect-nums[f] {
					k++
				} else {
					f--
				}
			}
		}
	}

	return ans
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(fmt.Sprintf("%v", fourSum(arr, -1)))
}
