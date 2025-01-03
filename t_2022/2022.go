package main

// https://leetcode.com/problems/convert-1d-array-into-2d-array/description/

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original)  {
		return make([][]int, 0)
	}

    matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		start := i * n
		end := start + n
		matrix[i] = original[start:end]               
	}

	return matrix
}

// if m*n != len(original)  {
// 	return make([][]int, 0)
// }

// matrix := make([][]int, m)

// k := 0

// for i := 0; i < m; i++ {
// 	matrix[i] = make([]int, n)
// 	for j := 0; j < n; j++ {
// 		if (k < len(original)) {
// 			matrix[i][j] = original[k]
// 			k++	
// 		}
// 	}                
// }

// return matrix


func main() {
	arr := []int{1,3,4,6,6,7}
	fmt.Println(construct2DArray(arr, 3,2))
}