package main

// https://leetcode.com/problems/move-zeroes/description/

import "fmt"

func moveZeroes(nums []int) {
	notZero := 0

	for i, v := range nums {
		if (v != 0 ) {
			nums[notZero]  = nums[i]
			notZero++
		}
	}

	for i := notZero; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	fmt.Println(arr)
	moveZeroes(arr)
	fmt.Println(arr)
}