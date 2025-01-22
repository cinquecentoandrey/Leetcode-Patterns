package main

import "fmt"

// https://leetcode.com/problems/palindrome-partitioning/description/


var res [][]string

func isPalindrome(s string, left int, rigth int) bool {
	for left < rigth {
		if s[left] != s[rigth] {
			return false
		}
		left++
		rigth--
	}
	return true
}

func traverse(idx int, part []string, s string) {
	fmt.Println(part, idx)
	if idx >= len(s) {
		copyCombination:= make([]string, len(part))
		copy(copyCombination, part)
		res = append(res, copyCombination)
		return 
	}
	
	for i := idx; i < len(s); i++ {
		if isPalindrome(s, idx, i) {
			part = append(part, s[idx:i+1])
			traverse(i+1, part, s)
			part = part[:len(part)-1]
		}
	}
}

func partition(s string) [][]string {
    res = [][]string{}
	traverse(0, make([]string, 0), s)
	return res
}

func main() {
	fmt.Println(partition("aab"))
}