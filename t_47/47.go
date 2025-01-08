package main

import "fmt"

// https://leetcode.com/problems/permutations-ii/

var res [][]int

// взяли мапу и по ней пошли переставлять
func traverse(perm []int, nums []int, countMap map[int]int) {
	if len(perm) == len(nums) {
		copyPerm := make([]int, len(perm))
   		copy(copyPerm, perm)
		res = append(res, copyPerm)
		return 
	}

	for k, _ := range countMap {
		if countMap[k] > 0 {
			countMap[k] -= 1
			perm = append(perm, k)

			traverse(perm, nums, countMap)
			
			perm = perm[:len(perm)-1]
			countMap[k] += 1
		}
	}

}

func permuteUnique(nums []int) [][]int {
    res = make([][]int, 0)
	countMap := make(map[int]int)

	for _, n := range nums {
		countMap[n] += 1
	}

	traverse(make([]int, 0), nums, countMap)
	return res
}

func main() {
	arr := []int{1,2,3}
	fmt.Println( permuteUnique(arr))
}


