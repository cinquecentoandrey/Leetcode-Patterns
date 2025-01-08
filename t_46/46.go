package main

import "fmt"

// https://leetcode.com/problems/permutations/description/

// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

var res [][]int

func traverse(current int, nums []int) {
	if current == len(nums) {
		copyNums := make([]int, len(nums))
   		copy(copyNums, nums)
		res = append(res, copyNums)
		return 
	}

	for i := current; i < len(nums); i++ {
		nums[i], nums[current] = nums[current], nums[i]

		traverse(current + 1, nums)

		nums[i], nums[current] = nums[current], nums[i]
	}
}

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	traverse(0, nums)
	return res
}

func main() {
	arr := []int{1,2,3}
	fmt.Println(permute(arr))
}