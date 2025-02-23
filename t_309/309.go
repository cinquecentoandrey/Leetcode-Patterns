package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/

type Key struct {
    i      int
    buying bool
}

func dfs(i int, buying bool, prices []int, dp map[Key]int) int {
	if i >= len(prices) {
		return 0
	}

	if val, ok := dp[Key{i, buying}]; ok {
		return val
	}

	var result int

	if buying {
		buy := dfs(i + 1, !buying, prices, dp) - prices[i]
		cooldown := dfs(i + 1, buying, prices, dp)
		result = max(buy, cooldown)
	} else {
		sell := dfs(i + 2, !buying, prices, dp) + prices[i]
		cooldown := dfs(i + 1, buying, prices, dp)
		result = max(sell, cooldown)
	}

	dp[Key{i, buying}] = result

	return result
}

func maxProfit(prices []int) int {
    dp := make(map[Key]int)

	return dfs(0, true, prices, dp)
}