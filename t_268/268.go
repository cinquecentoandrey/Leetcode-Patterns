package main

// https://leetcode.com/problems/missing-number/description/

import "fmt"

func main() {
	arr := [3]int{1,2,3}
	fmt.Println(missingNumber(arr[:]))
}

func missingNumber(nums []int) int {
	idxArr := make([]bool, len(nums)+1)

	for _, v := range nums {
		idxArr[v] = true
	}

	for i, v := range idxArr {
		if !v {
			return i
		}
	}

	return int(-1)
}