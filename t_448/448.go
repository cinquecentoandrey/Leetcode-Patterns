package main

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

import "fmt"

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(arr))
}

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	idxArr := make([]bool, len(nums)+1)

	for _, v := range nums {
		idxArr[v] = true
	}

	for i := 1; i < len(idxArr); i++ {
		if !idxArr[i] {
			res = append(res, i)
		}
	}

	return res
}