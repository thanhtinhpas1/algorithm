package main

import (
	"fmt"
)

const MaxUint = ^uint32(0)
const MaxInt = int64(MaxUint >> 1)
const MinInt = -MaxInt - 1

// convert string to int
// - ignore any whitespace
// - determine the sign with '-' or '+'
func myAtoi(s string) int {
	s = trimLeadingSpace(s)

	if len(s) == 0 {
		return 0
	}

	var sign int8
	var result int64

	arr := []rune(s)
	if arr[0] == '-' {
		sign = -1
		arr = arr[1:]
	} else if arr[0] == '+' {
		sign = 1
		arr = arr[1:]
	} else {
		sign = 1
	}

	if len(arr) == 0 {
		return 0
	}

	if arr[0]-'0' < 0 || arr[0]-'0' > 9 {
		return 0
	}

	result = int64(arr[0] - '0')
	arr = arr[1:]

	for _, n := range arr {
		newNum := int64(n - '0')
		if newNum > 9 || newNum < 0 {
			break
		}

		result *= 10
		result += newNum

		if result*int64(sign) > MaxInt {
			return int(MaxInt)
		}

		if result*int64(sign) < MinInt {
			return int(MinInt)
		}
	}

	return int(result * int64(sign))
}

func trimLeadingSpace(s string) string {
	var firstChar int
	arr := []rune(s)
	for i, _ := range arr {
		if arr[i] != ' ' {
			firstChar = i
			break
		}
	}

	return s[firstChar:]
}

func main() {
	fmt.Println(myAtoi("   +0 123"))
}
