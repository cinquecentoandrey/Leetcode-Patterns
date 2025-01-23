package main

import "fmt"

// https://leetcode.com/problems/house-robber-ii/description/

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp1 := make([]int, n + 1)
	dp2 := make([]int, n + 1)
	
	for i := 2; i< n+1; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2] + nums[i-2])
	}

	nums = nums[1:]
	for i := 2; i < n+1; i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2] + nums[i-2])
	}

	return max(dp1[n], dp2[n])
}	

func main() {
	fmt.Println(rob([]int{1,2,3}))
	fmt.Println(rob([]int{1,2,3,1}))
}