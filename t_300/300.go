package main

import "fmt"

// https://leetcode.com/problems/longest-increasing-subsequence/description/

func lengthOfLIS(nums []int) int {
    lis := make([]int, len(nums))

	for i := 0; i < len(lis); i++ {
		lis[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println(nums[i])
		for j := i + 1;  j < len(nums); j++ {
			if nums[i] < nums[j] {
				fmt.Println(lis)
				lis[i] = max(lis[i], 1 + lis[j])
				
			}
		}
		
		fmt.Println()
	}

	res := 0 
	
	for _, v := range lis {
		if v  > res {
			res = v
		}
	}

	return res
}

// func lengthOfLIS(nums []int) int {
//     maxLen := 0
// 	dp := make([]int, len(nums))

// 	for i := 1; i <= len(nums) - 1; i++ {
// 		for j := 0; j <= i  - 1; j++ {
// 			if nums[j] < nums[i] {
// 				dp[i] = max(dp[i], dp[j] + 1)
// 			}
// 		}
// 	} 
// 	fmt.Println(dp)
// 	for _, v := range dp {
// 		if v > maxLen {
// 			maxLen = v
// 		}
// 	}

// 	return maxLen
// }

func main() {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
}