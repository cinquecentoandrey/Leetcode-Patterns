package main

import "fmt"

// https://leetcode.com/problems/subsets/

var res [][]int

func traverse(i int, subset []int, nums []int) {
	if i == len(nums) {
		copySubset:= make([]int, len(subset))
		copy(copySubset, subset)
		res = append(res, copySubset)
		return 
	}

	subset = append(subset, nums[i])
	traverse(i + 1, subset, nums)
	subset = subset[:len(subset)-1]
	traverse(i + 1, subset, nums)
}

func subsets(nums []int) [][]int {
    res = make([][]int, 0)
	traverse(0, make([]int, 0), nums)
	return res
}



// прикол
// func subsets(nums []int) [][]int {
// 	res := make([][]int, 1)

// 	for _, v := range nums {
// 		newSubsets := make([][]int, 0)
// 		for _, subset := range res {
// 			newSubset := append([]int{}, subset...)
// 			newSubset = append(newSubset, v)
// 			newSubsets = append(newSubsets, newSubset)
// 		}
// 		res = append(res, newSubsets...)
// 	}

// 	return res
// }

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(subsets(arr))
}
