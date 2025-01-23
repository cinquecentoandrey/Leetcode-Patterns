package main

import "fmt"

// https://leetcode.com/problems/house-robber/


func rob(nums []int) int {
	dp := make([]int, len(nums) + 2)
	
	for i := 2; i < len(nums) + 2; i++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i-2])
	}

	return dp[len(nums)+1]
}	


// func rob(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

//     res := make([]int, len(nums))
// 	res[0] = nums[0]
// 	res[1] = max(res[0], nums[1]) 
// 	fmt.Println(res)
// 	for i := 2; i < len(nums); i++ {
// 		res[i] = max(res[i-1], res[i-2] + nums[i])
// 	}

// 	return res[len(res) - 1]
// }	

func main() {
	fmt.Println(rob([]int{2,7,9,3,1}))
	
	//fmt.Println(rob([]int{2,1,3,1}))
}