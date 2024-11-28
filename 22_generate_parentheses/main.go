package main

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(l, r int, s string)

	dfs = func(l int, r int, s string) {
		if len(s) == n*2 {
			ans = append(ans, s)
			return
		}

		if l < n {
			dfs(l+1, r, "("+s)
		}

		if r < l {
			dfs(l, r+1, s+")")
		}
	}

	dfs(0, 0, "")

	return ans
}

func main() {
	fmt.Println(fmt.Sprintf("%v", generateParenthesis(3)))
}
