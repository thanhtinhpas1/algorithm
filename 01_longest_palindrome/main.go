package main

import (
	"fmt"
	"math"
	"strings"
)

// The simplest way is brute force, which is pick all possible starting positions then loop through them
// to check whether whether it is palindrome, and if it is longest, mark it as the result.

/**
Approach:
We recognize that palindrome mirros it around center. There are 2n - 1 such center positions.
- Each position can be a center
- Two of each position can be a center: 2 * n/2 (on the left and on the right)
**/

// Complexity is O(n^2)
// Space complexity is O(n^2)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var (
		arr  = []rune(s)
		left int
		max  []rune
	)

	// we go through from left to right
	for i := left + 1; i < len(arr); i++ {
		space := i - left + 1

		// with event space, we will consider the current item & next to will be center
		if space%2 == 0 {
			midLeft := left + (space / 2) - 1
			midRight := midLeft + 1
			isPalin := true
			for j := 0; j < space/2; j++ {
				if arr[midLeft-j] != arr[midRight+j] {
					isPalin = false
					break
				}
			}

			if isPalin {
				if space > len(max) {
					max = arr[left : i+1]
				}
			}
		} else { // with odd space, we will consider the current item as center
			mid := left + (space / 2)

			isPalin := true
			for j := 0; j < space/2; j++ {
				if arr[mid-j] != arr[mid+j] {
					isPalin = false
					break
				}
			}

			if isPalin && space > len(max) {
				max = arr[left : i+1]
			}
		}
	}

	return string(max)
}

func expandFromCenter(arr []rune, left, right int) []rune {
	if len(arr) <= 1 {
		return arr
	}

	for left >= 0 && right < len(arr) && arr[left] == arr[right] {
		left--
		right++
	}

	return arr[left+1 : right]
}

// Time complexity O(n^2)
// Space complexity O(1)
func longestPalindrome_02(s string) string {
	var (
		arr = []rune(s)
		max []rune
	)

	for i := 0; i < len(arr)-1; i++ {
		even := expandFromCenter(arr, i, i)
		odd := expandFromCenter(arr, i, i+1)

		if len(even) > len(max) {
			max = even
		}

		if len(odd) > len(max) {
			max = odd
		}
	}

	return string(max)
}

// Time complexity O(n^2)
// Space complexity O(n^2)
func longestPalindrome_03(s string) string {
	var (
		arr    = []rune(s)
		maxStr = []rune{arr[0]}
		dp     = make(map[int][]bool)
	)

	// init dp with all false values
	for idx := range arr {
		dp[idx] = make([]bool, len(arr))
	}

	for i := range arr {
		// center as palindrome
		dp[i][i] = true

		for j := 0; j < i; j++ {
			if arr[i] == arr[j] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true

				if i-j+1 > len(maxStr) {
					maxStr = arr[j : i+1]
				}
			}
		}
	}

	return string(maxStr)
}

// Time complexity: O(N)
// Space complexity: O(1)
// Using manacher's algorithm
// We maintain a boundary palindrome center, we use previous result to determine whether can expand the
// palindrome characters or not
func longestPalindrome_04(s string) string {
	var (
		arr    = []rune(s)
		maxStr = []rune{arr[0]}
		right  = 0
		center = 0
		max    = 1
	)

	arr = []rune(fmt.Sprintf("#%s#", strings.ReplaceAll(s, "", "#"))) // make the string no odd or even case
	dp := make([]int, len(arr))                                       // result of the item can expand as palindrome

	for i := 0; i < len(arr); i++ {
		if i < right {
			dp[i] = int(math.Min(float64(right-i), float64(dp[2*center-i]))) // compare from right - i or previous result
		}

		for i-dp[i]-1 >= 0 && i+dp[i]+1 < len(arr) && arr[i-dp[i]-1] == arr[i+dp[i]+1] {
			dp[i]++
		}

		if i+dp[i] > right {
			center = i
			right = i + dp[i]
		}

		if dp[i] > max {
			max = dp[i]
			maxStr = arr[i-dp[i] : i+dp[i]+1]
		}
	}

	return strings.ReplaceAll(string(maxStr), "#", "")
}

func main() {
	fmt.Println(longestPalindrome("abbacadef"))
	fmt.Println(longestPalindrome_02("abbacadef")) // this is fatest solution you can think about when in the interview session
	fmt.Println(longestPalindrome_03("abbacadef"))
	fmt.Println(longestPalindrome_04("abbacadef")) // this is the best solution you can think about when you having more time
}
