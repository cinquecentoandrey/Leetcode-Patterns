package main

// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/description/

func backtracking(i int, k int, subsetSum int, target int, nums []int, used []bool) bool {
	if k == 0 {
		return true
	}

	if subsetSum == target {
		return backtracking(0, k - 1, 0, target, nums, used)
	}

	for j := i; j < len(nums); j++ {
		if used[j] || subsetSum  + nums[j] > target {
			continue
		}
		used[j] = true
		if backtracking(j + 1, k, subsetSum + nums[j], target, nums, used) {
			return true
		}
		used[j] = false
	}
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
    var totalSum int
    for _, v := range nums {
        totalSum += v
    }

	target := totalSum / k

	used := make([]bool, len(nums))

	return backtracking(0, k, 0, target, nums, used)
}