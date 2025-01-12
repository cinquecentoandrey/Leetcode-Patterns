package main

// https://leetcode.com/problems/combination-sum-ii/

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
	traverse(index+1, total + candidates[index], currentCumobination, target, candidates)
	currentCumobination = currentCumobination[:len(currentCumobination)-1]
	for index + 1 < len(candidates) && candidates[index] == candidates[index + 1] {
		index++
	}

	traverse(index + 1, total, currentCumobination, target, candidates)

}

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	candidates = quickSort(candidates)
	traverse(0, 0, make([]int, 0), target, candidates)
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

