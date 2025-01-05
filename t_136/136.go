package main

// https://leetcode.com/problems/single-number/

// тут все просто - xor
func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}

	return res
}