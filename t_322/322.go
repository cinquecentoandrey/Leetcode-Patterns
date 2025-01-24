package main

import "fmt"

// https://leetcode.com/problems/coin-change/description/

// вроде понял прикол, но тяжко мне даются dp проблемы...
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)

	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for a := 1; a < amount+1; a++ {
		for _, c := range coins {
			if a - c >= 0 {
				dp[a] = min(dp[a], 1 + dp[a-c])
			}
		}
	}

	if dp[amount] != amount + 1 {
		return dp[amount]
	} else {
		return -1
	}
}

func main() {
	fmt.Println(coinChange([]int{1,2,5}, 11))
}