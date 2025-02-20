package main

import "fmt"

/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
	}

	pow := float64(1)
	num := int(abs(float64(n)))

	for num > 0 {
		pow *= x
		num--
	}

	return pow
}

func abs(x float64) float64 {
	if x < 0 {
		return x * -1
	}
	return x
}

/*
(11)10 is (1011)2
1   0   1   1
2^3 2^2 2^1 2^0

OR we can also write this as
1 0 1 1
8 4 2 1

2 ^ (8 + 2 + 1) = 2 ^ 8 * 2 ^ 2 * 2 ^ 1

Complexity: O(log n)
*/
func myPow_2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
	}

	num := int(abs(float64(n)))
	pow := float64(1)

	for num != 0 {
		if (num & 1) != 0 {
			pow *= x
		}

		x *= x
		num >>= 1
	}

	return pow
}

func main() {
	x := 2.0
	n := -2

	fmt.Printf("%v", myPow(x, n))   // expect: 1024.00000
	fmt.Printf("%v", myPow_2(x, n)) // expect: 1024.00000
}
