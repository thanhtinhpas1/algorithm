package main

import "fmt"

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}

	arr := []rune(s)
	result := []rune{arr[0]}

	var i int
	var duplicate int
	for i < len(arr)-1 {
		if arr[i] != arr[i+1] {
			result = append(result, arr[i+1])
			duplicate = 0
		}

		if arr[i] == arr[i+1] {
			duplicate++

			if duplicate < 2 {
				result = append(result, arr[i+1])
			}
		}

		i++
	}

	return string(result)
}

func main() {
	fmt.Println(makeFancyString("leeetcode"))
	fmt.Println(makeFancyString("aaabaaaa"))
}
