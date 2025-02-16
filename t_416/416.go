package main

// https://leetcode.com/problems/partition-equal-subset-sum/description/

// медленный бэктрэк 
func backtracking(nums []int, index int, target int, memo map[[2]int]bool) bool {
	if target == 0 {
		return true
	}

	if index < 0 || target < 0 {
		return false
	}

	key := [2]int{index, target}
    if val, found := memo[key]; found {
        return val
    }

	take := backtracking(nums, index - 1, target - nums[index], memo)
	notTake := backtracking(nums, index-1, target, memo)

	memo[key] = take || notTake
	return memo[key]
}

func canPartition(nums []int) bool {
	var totalSum int
    for _, v := range nums {
        totalSum += v
    }

    if totalSum % 2 != 0 {
        return false
    }

	target := totalSum / 2
	memo := make(map[[2]int]bool)
	return backtracking(nums, len(nums) - 1, target, memo)
}


// жестко наговнякал
// func canPartition(nums []int) bool {
//     var totalSum int
//     for _, v := range nums {
//         totalSum += v
//     }

//     if totalSum % 2 != 0 {
//         return false
//     }

//     target := totalSum / 2
//     dp := make(map[int]bool)
//     dp[0] = true
//     find := false
//     for i := len(nums) - 1; i > -1; i-- {
// 		newElements := make(map[int]bool)

//         for t := range dp {
//             if t+nums[i] <= target {
//                 newElements[t+nums[i]] = true   
//             }
//             if t+nums[i] == target {
//                 find = true
//             }
// 			newElements[t] = true
//         }

//         dp = newElements
//         if find {
//             return true
//         }
// 	}

//     return dp[target]
// }