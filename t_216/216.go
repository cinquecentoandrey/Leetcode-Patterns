package main

import (
	"fmt"
)

// https://leetcode.com/problems/combination-sum-iii/

var res [][]int

func traverse(index int, total int, k int, currentCumobination []int, target int, candidates []int) {
	if target == total && k == len(currentCumobination){
		copyCombination:= make([]int, len(currentCumobination))
		copy(copyCombination, currentCumobination)
		res = append(res, copyCombination)
		return 
	}

	if index >= len(candidates) || total > target || len(currentCumobination) > k{
		return
	}

	for i := index; i < len(candidates); i++ {
		currentCumobination = append(currentCumobination, candidates[i])
		traverse(i + 1, total+candidates[i], k, currentCumobination, target, candidates)
		currentCumobination = currentCumobination[:len(currentCumobination)-1]
	}

}

func combinationSum3(k int, n int) [][]int {
    nums := []int{1,2,3,4,5,6,7,8,9}
	res = [][]int{}

	traverse(0, 0, k, []int{}, n, nums)

	return res
}

func main() {
	fmt.Println(combinationSum3(4, 1))
}