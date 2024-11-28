package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	isNegative := (dividend < 0) != (divisor < 0)
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	result := 0

	for dividend >= divisor {
		tmpDivisor := divisor
		multiple := 1

		for dividend >= (tmpDivisor << 1) {
			tmpDivisor <<= 1
			multiple <<= 1
		}

		dividend -= tmpDivisor
		result += multiple
	}

	if isNegative {
		result = -result
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}

func main() {
	fmt.Println(divide(-2147483648, -1))
}
