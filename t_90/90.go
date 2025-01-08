package main

import "fmt"

// https://leetcode.com/problems/subsets-ii/

var res [][]int

func backtrack(i int, subset []int, nums []int) {
	if i == len(nums) {
		copySubset:= make([]int, len(subset))
		copy(copySubset, subset)
		res = append(res, copySubset)
		return 
	}

	subset = append(subset, nums[i])
	backtrack(i + 1, subset, nums)

	subset = subset[:len(subset)-1]
	for i + 1 < len(nums) && nums[i] == nums[i + 1] {
		i++
	}

	backtrack(i + 1, subset, nums)
}

func subsetsWithDup(nums []int) [][]int {
	nums = quickSort(nums)
	backtrack(0, make([]int, 0), nums)
	return res
}


func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr 
    }

    pivot := arr[len(arr)/2] 
    var left []int
    var right []int

    for i, val := range arr {
        if i == len(arr)/2 {
            continue 
        }
        if val <= pivot {
            left = append(left, val) 
        } else {
            right = append(right, val) 
        }
    }

    return append(append(quickSort(left), pivot), quickSort(right)...)
}

// по рабоче крестъянски отсортировали и забрутфорсили
// func checkForDup(set [][]int, subset []int) bool {
//     for i := 0; i < len(set); i++ {
//         if len(set[i]) == len(subset) { 
//             match := true
//             for j := 0; j < len(subset); j++ {
//                 if subset[j] != set[i][j] { 
//                     match = false
//                     break
//                 }
//             }
//             if match {
//                 return true 
//             }
//         }
//     }

//     return false 
// }

// func subsetsWithDup(nums []int) [][]int {
// 	nums = quickSort(nums)
//     res := make([][]int, 1)

// 	for _, v := range nums {
// 		newSubsets := make([][]int, 0)
// 		for _, subset := range res {
// 			newSubset := append([]int{}, subset...)
// 			newSubset = append(newSubset, v)
// 			if !checkForDup(res, newSubset) {
// 				newSubsets = append(newSubsets, newSubset)
// 			}
// 		}
// 		res = append(res, newSubsets...)
// 	}

// 	return res
// }

func main() {
	arr := []int{0}
	fmt.Println(subsetsWithDup(arr))
}