package main

import "fmt"

// https://leetcode.com/problems/squares-of-a-sorted-array/description/

// func sortedSquares(nums []int) []int {

// 	left := 0
// 	right := len(nums) - 1

// 	for i, v := range nums {
// 		nums[i] = v * v
// 	}

// 	for left <= right {
// 		if nums[left] > nums[right] {
// 			tmp := nums[left]
// 			nums[left] = nums[right]
// 			nums[right] = tmp
// 		}
// 		right--
// 	}

// 	return nums
// }

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, len(nums))

	for i := len(res) - 1; i >=0; i-- {
		if nums[left] * nums[left] > nums[right] * nums[right] {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
	}

	return res
}

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(arr))
}