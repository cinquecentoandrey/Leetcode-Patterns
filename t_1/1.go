package main

// https://leetcode.com/problems/two-sum/

import "fmt"

func twoSum(nums []int, target int) []int {
	cMap := make(map[int]int)
	cMap[nums[0]] = 0

	for i := 1; i < len(nums); i++ {
		elem := target - nums[i]
		value, exists := cMap[elem]
		if !exists {
			cMap[nums[i]] = i
		} else {
			return []int{value, i}
		}
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}