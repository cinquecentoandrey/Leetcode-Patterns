package main

import "fmt"

// https://leetcode.com/problems/target-sum/description/

// чет для умников
func findTargetSumWays(nums []int, target int) int {
	
	dp := make(map[int]int, 0)
	dp[0] = 1

	for _, v := range nums {
		nxt_dp := make(map[int]int, 0)
		for cur, count := range dp {
			nxt_dp[cur + v] += count
			nxt_dp[cur - v] += count
		}
		dp = nxt_dp
	}

	return dp[target]
}


// ниавное
// var counter int

// func traverse(nums []int, target int, index int, total int){
// 	if index == len(nums) {
// 		if total == target {
// 			counter++
// 			return 
// 		} else {
// 			return 
// 		}
// 	}

// 	traverse(nums, target, index+1, total + nums[index]) 
// 	traverse(nums, target, index+1, total - nums[index])
// }

// func findTargetSumWays(nums []int, target int) int {
//     counter = 0
// 	traverse(nums, target, 0, 0)
// 	return counter
// }

func main() {
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1}, 3))
}