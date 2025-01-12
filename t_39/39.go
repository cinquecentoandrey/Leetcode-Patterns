package main

import "fmt"

// https://leetcode.com/problems/combination-sum/

var res [][]int

func traverse(index int, total int, currentCumobination []int, target int, candidates []int) {

	if target == total {
		copyCombination:= make([]int, len(currentCumobination))
		copy(copyCombination, currentCumobination)
		res = append(res, copyCombination)
		return 
	}

	if index >= len(candidates) || total > target {
		return
	}

	currentCumobination = append(currentCumobination, candidates[index])
	traverse(index, total + candidates[index], currentCumobination, target, candidates)
	currentCumobination = currentCumobination[:len(currentCumobination)-1]
	traverse(index + 1, total, currentCumobination, target, candidates)
}

func combinationSum(candidates []int, target int) [][]int {
    res = [][]int{}
	traverse(0, 0, make([]int, 0), target, candidates)
	return res
}

func main() {
	arr := []int{2,3,5}
	fmt.Println(combinationSum(arr, 7))
}