package main

// https://leetcode.com/problems/jump-game/

func canJump(nums []int) bool {
    goal := len(nums) - 1

	for i := goal; i > 0; i-- {
		if i + nums[i] >= goal {
			goal = i
		}
	}

	if goal == 0 {
		return true
	} else {
		return false
	}
}