package main

import (
	"fmt"
)

// https://leetcode.com/problems/letter-case-permutation/description/

// минут 10 решение заняло, все интуитивно
// A - Z [65 - 90]
// a - z [97 - 122]
// 0 - 9 [48 - 57]
func traverse(chars []byte, currentIndex int, permutations *[]string) {
	if currentIndex >= len(chars) {
		*permutations = append(*permutations, string(chars))
		return
	}

	if chars[currentIndex] >= 65 && chars[currentIndex] <= 90 {
		chars[currentIndex] += 32
		traverse(chars, currentIndex + 1, permutations)

		chars[currentIndex] -= 32
		traverse(chars, currentIndex + 1, permutations)
	} else if chars[currentIndex] >= 97 && chars[currentIndex] <= 122 { 
		chars[currentIndex] -= 32
		traverse(chars, currentIndex + 1, permutations)

		chars[currentIndex] += 32
		traverse(chars, currentIndex + 1, permutations)
	} else if chars[currentIndex] >= 48 && chars[currentIndex] <= 57 {
		traverse(chars, currentIndex + 1, permutations)
	}

}

func letterCasePermutation(s string) []string {
    var permutations []string
	traverse([]byte(s), 0, &permutations)
	return permutations
}

func main() {
 	str := "3z4" 
	//str := "a1b2"
	res := letterCasePermutation(str)
	fmt.Println(res)
}