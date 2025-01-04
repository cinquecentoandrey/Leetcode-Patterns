package main

// https://leetcode.com/problems/find-all-duplicates-in-an-array/

// достойное решение от андрюхи
func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	
	return num
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

    for i := 0; i < len(nums); i++ {
		val := abs(nums[i])
		if nums[val - 1] < 0 {
			res = append(res, val)
		} else {
			nums[val - 1] *=-1
		}
	}

	return res
}

// херота от андрюхи (ну можно используя массив булов но не удовл условию)

// func findDuplicates(nums []int) []int {
// 	bools := make([]bool, len(nums) + 1)
// 	res := make([]int, 0)
//     for i := 0; i < len(nums); i++ {
// 		if (bools[nums[i]]) {
// 			res = append(res, nums[i])
// 		} 
// 		bools[nums[i]] = true
// 	}

// 	return res
// }
