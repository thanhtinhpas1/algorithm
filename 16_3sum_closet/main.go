package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// find three integers such that the sum is closest to target
	// by anyway we should kee the minimum of sum with the target

	ans := int(^uint(0) >> 1)
	minAbs := int(^uint(0) >> 1)

	sort.Ints(nums) // O(NLog(N))
	//-4 -1 1 2
	// 1 + 4 + 1 = 6, c = 6

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1

		tmpTarget := target - nums[i]
		for j < k {
			tmpSum := nums[j] + nums[k]

			if tmpAbs := abs(target - (tmpSum + nums[i])); tmpAbs < minAbs {
				minAbs = tmpAbs
				ans = tmpSum + nums[i]
			}

			if tmpSum < tmpTarget {
				j++
			} else {
				k--
			}
		}
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}

func main() {
	arr := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	target := -2
	fmt.Println(threeSumClosest(arr, target))
}
