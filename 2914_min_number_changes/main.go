package main

import (
	"fmt"
)

func minChanges(s string) int {
	arr := []rune(s)
	count := 0
	
	for i := 0; i < len(s); i += 2 {
		if !checkPerfect([]rune{arr[i], arr[i+1]}) {
			count++
		}
	}

	return count
}

func checkPerfect(arr []rune) bool {
	if arr[0] != arr[1] {
		return false
	}

	return true
}

func main() {
	fmt.Println(minChanges("0110"))
}