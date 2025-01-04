package main

// https://leetcode.com/problems/product-of-array-except-self/description/

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res:= make([]int, n)
	left := 1
	right := 1

	for i := 0; i < n; i++ {
		res[i] = left
		left *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}

// рабоче крестьянский
// func productExceptSelf(nums []int) []int {
//     n := len(nums)
// 	left := make([]int, n)
// 	right := make([]int, n)
// 	res:= make([]int, n)
// 	left[0] = 1
// 	right[n-1] = 1

// 	for i := 1; i < n; i++ {
// 		left[i] = left[i - 1] * nums[i-1]
// 	}

// 	for i := n - 2; i >= 0; i-- {
// 		right[i] = right[i + 1] * nums[i + 1]
// 	}

// 	for i := 0; i < n; i++ {
// 		res[i] = left[i]*right[i]
// 	}

// 	return res
// }

func main() {
	productExceptSelf([]int{1,2,3,4})
	productExceptSelf([]int{-1,1,0,-3,3})
}