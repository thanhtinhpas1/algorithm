package main

import "fmt"

func solution(n, k int) int {
	if n <= 1 {
		return n
	}

	count := 0
	remain := n

	// we will calculate how many times we should double price
	// first time: 8 / 2 = 4 (remain)
	// second time: 4 / 2 = 2 (remain)
	// ...
	for i := 0; i < k && remain > 1; i++ {
		remain = remain / 2
		count++
	}

	// with remain, we will need remain - 1 times to reach out
	count += remain - 1

	return count
}

func main() {
	fmt.Printf("%v", solution(8, 1))
}
