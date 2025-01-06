package main

// https://leetcode.com/problems/set-matrix-zeroes/description/

import "fmt"

// есть способ с решением без доп памяти, но там такая каша, что делать этого не хочется
func setZeroes(matrix [][]int) {
	zeroIndexes := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeroIndexes = append(zeroIndexes, i, j)
			}
		}
	}

	for len(zeroIndexes) > 0 {
		rowIdx := zeroIndexes[0]
		colIdx := zeroIndexes[1]

		for i := 0; i < len(matrix[rowIdx]); i++ {
			matrix[rowIdx][i] = 0
		}

		for i := 0; i < len(matrix); i++ {
			matrix[i][colIdx] = 0
		}
		zeroIndexes = zeroIndexes[2:]
	}
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

}

func main() {

	matrix3x3 := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	setZeroes(matrix3x3)
	fmt.Println()

	matrix3x4 := [][]int{
		{0, 2, 3, 0},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	setZeroes(matrix3x4)
}