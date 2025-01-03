package main

// https://leetcode.com/problems/maximum-average-subarray-i/description/

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}
	
	curSum := maxSum
	for i := k; i < len(nums); i++ {
		curSum += nums[i] - nums[i-k]
		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}