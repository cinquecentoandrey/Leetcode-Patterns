package main

// https://leetcode.com/problems/is-subsequence/description/

import "fmt"

func isSubsequence(s string, t string) bool {
	left := 0
	right := 0

	for left < len(s) && right < len(t) {
		if (s[left] == t[right]) {
			left++
			right++
		} else {
			right++
		}
	}

	return left == len(s)
}

func main() {
	s := "abc"
	t := "ah1gdc"

	fmt.Println(isSubsequence(s, t))
}