package main

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

import "fmt"

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(arr))
}

// func findDisappearedNumbers(nums []int) []int {
// 	res := make([]int, 0)
	
// 	i := 0
// 	for i < len(nums) {
// 		pos := nums[i] - 1
// 		if nums[i] != nums[pos] {
// 			nums[i], nums[pos] = nums[pos], nums[i]
// 		} else {
// 			i++
// 		}
// 	}

// 	for i, v := range nums {
// 		if v != i+1 {
// 			res = append(res, i+1)
// 		}
// 	}

// 	return res
// }

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