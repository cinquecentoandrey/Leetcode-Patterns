package main

import "fmt"

// https://leetcode.com/problems/generate-parentheses/

var res []string 

func traverse(open int, close int, path string, n int) {
	if open == close && open == n && close == n {
		res = append(res, path)
		return
	}

	if open < n {
		traverse(open+1, close, path + "(", n)
	}

	if close < open {
		traverse(open, close+1, path + ")", n)
	}
}

func generateParenthesis(n int) []string {
    res = make([]string, 0)
	path := ""
	traverse(0, 0, path, n)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}