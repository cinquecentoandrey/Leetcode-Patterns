package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}

func maxProfit(prices []int) int {
	minS := prices[0]
	maxS := 0

	for i := 0; i < len(prices); i++ {
		minS = min(prices[i], minS)
		diff := (prices[i] - minS)
		maxS = max(diff, maxS)
		fmt.Printf("Min: %d, max: %d\n", minS, maxS)
	}

	return maxS
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// ну ктож виноватъ что в го нет обычной тернарки и эта хрень долго выполняется
// func Ternary[T any](cond bool, vtrue T, vfalse T) T {
// 	if cond {
// 		return vtrue
// 	}
// 	return vfalse
// }