package main

import "fmt"

// https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	curSum := 0
	maxSub := nums[0]
	for i := 0; i < len(nums); i++ {
		if curSum < 0 {
			curSum = 0
		}
		curSum += nums[i]
		maxSub = max(curSum, maxSub)

	}

	return maxSub
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}