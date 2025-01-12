package main

import "fmt"

// https://leetcode.com/problems/combinations/

var res [][]int

func traverse(start int, k int, combination []int, nums []int) {
	if len(combination) == k {
		copyCombination := make([]int, k)
		copy(copyCombination, combination)
		res = append(res, copyCombination)
		return
	}

	for i := start; i < len(nums); i++ {
		combination = append(combination, nums[i])
		traverse(i + 1, k , combination, nums)
		combination = combination[:len(combination)-1]
	}
}

func combine(n int, k int) [][]int {
    nums := make([]int, n)
	res = [][]int{}
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	traverse(0, k, []int{}, nums)

	return res
}


func main() {
	combine(4, 2)
}