package main

import "fmt"

/*
given two non-negative numbers num1 and num2 represented as strings, return the product of `num1` and `num2`,
also represent in string
Note: not using built in Big Integer library or convert the inputs to integer directly
*/
func multiply(i1 string, i2 string) string {
	/**
	Let's look at an example. Consider 123∗456, it can be written as,
	⟹(123 ∗ (6 + 50 + 400))
	⟹(123 ∗ 6) + (123 ∗ 50) + (123 ∗ 400)
	⟹(123 ∗ 6) + (123 ∗ 5 ∗ 10)+(123 ∗ 4 ∗ 100)
	*/

	if i1 == "0" || i2 == "0" {
		return "0"
	}

	num1 := reverse(i1)
	num2 := reverse(i2)

	var result [][]rune
	maxLength := 0

	for i := 0; i < len(num1); i++ {
		carry := rune(0)
		line := make([]rune, i)
		for j := 0; j < len(num2); j++ {
			multiply := (num1[i] - '0') * (num2[j] - '0')
			line = append(line, (multiply+carry)%10)
			carry = (multiply + carry) / 10
		}

		if carry > rune(0) {
			line = append(line, carry)
		}
		if len(line) > maxLength {
			maxLength = len(line)
		}

		result = append(result, line)
	}

	for i := 0; i < len(result); i++ {
		for len(result[i]) < maxLength {
			result[i] = append(result[i], rune(0))
		}
	}

	var multiply []rune
	carry := rune(0)
	for i := 0; i < maxLength; i++ {
		sum := rune(0)
		for j := 0; j < len(result); j++ {
			sum += result[j][i]
		}

		multiply = append(multiply, (sum+carry)%10)
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		multiply = append(multiply, carry)
	}

	multiply = reverse(string(multiply))
	for i := 0; i < len(multiply); i++ {
		multiply[i] += '0'
	}

	return string(multiply)
}

func reverse(num string) []rune {
	i := 0
	j := len(num) - 1
	arr := []rune(num)
	for i <= j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	return arr
}

func main() {
	rs := multiply("456", "0")
	fmt.Println(rs)
}
