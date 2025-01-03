package main

// https://leetcode.com/problems/binary-search/description/

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for left < right {
		mid = (left + right) / 2
		
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{5}
	fmt.Println(search(arr, 5))
}