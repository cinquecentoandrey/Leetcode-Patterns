package main

import "fmt"

// https://leetcode.com/problems/combination-sum-iv/description/

// что то похожее на фибоначи, смотрим текущий тотал и смотрим предыдущю сумму
func combinationSum4(nums []int, target int) int {
    dp := make([]int, target+1)
    dp[0] = 1

    for total := range dp {
		fmt.Println("total: ", total)
        for j := 0; j < len(nums); j++ {
            if total >= nums[j] {
				fmt.Println("\tnums[j]: ", nums[j])
                dp[total] += dp[total - nums[j]]
				fmt.Println("\tdp: ", dp)
            }
        }
    }

    return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1,2,3}, 4))
}


// func getOrDefault(m map[int]int, key int, defaultValue int) int {
//     if val, exists := m[key]; exists {
//         return val
//     }
//     return defaultValue
// }

// func combinationSum4(nums []int, target int) int {
//     dp := make(map[int]int)
// 	dp[0] = 1

// 	for total := 1; total < target +1; total++ {
// 		dp[total] = 0
// 		for _, n := range nums {
// 			dp[total] += getOrDefault(dp, total - n, 0)
// 		}
// 	}
// 	return dp[target]
// }