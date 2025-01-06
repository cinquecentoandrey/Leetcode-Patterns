package main

import (
	"fmt"
	//"time"
	//"sort"
)

// https://leetcode.com/problems/longest-consecutive-sequence/submissions/1500004027/

// хз, вроде o(n)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longestConsecutive := 1
	tmplongestConsecutive := 1

	numMap := make(map[int]bool)

	for _, v := range nums {
		numMap[v] = true
	}

	for num := range numMap {
		if !numMap[num - 1] {
			for numMap[num + tmplongestConsecutive] {
				tmplongestConsecutive++
			}
			if tmplongestConsecutive > longestConsecutive {
				longestConsecutive = tmplongestConsecutive
			} 
			tmplongestConsecutive = 1
		}
	}

	return longestConsecutive
}

// o(n*log(n))
// func longestConsecutive(nums []int) int {
// 	longestConsecutive := 1
// 	tmplongestConsecutive := longestConsecutive

// 	sort.Ints(nums)
// 	fmt.Println(nums)

// 	for i := 0; i < len(nums) - 1; i++ {
// 		if nums[i+1] - nums[i] == 1 {
// 			tmplongestConsecutive++
// 		} else if nums[i+1] - nums[i] == 0 {
// 			continue
// 		} else if nums[i+1] - nums[i] > 1 {
// 			if tmplongestConsecutive > longestConsecutive {
// 				longestConsecutive = tmplongestConsecutive
// 				tmplongestConsecutive = 0
// 			} else {
// 				tmplongestConsecutive = 0
// 			}
// 		}
// 	}

// 	if tmplongestConsecutive > longestConsecutive {
// 		longestConsecutive = tmplongestConsecutive
// 		tmplongestConsecutive = 0
// 	} else {
// 		tmplongestConsecutive = 0
// 	}

// 	return longestConsecutive
// }

// fmt.Printf("\tConseq: %d, currentVal: %d, nums[%d]: %d\n", tmplongestConsecutive, currentVal, i, nums[i])

func main() {
	arr := []int{100,4,200,1,3,2}
	fmt.Println(longestConsecutive(arr))
}