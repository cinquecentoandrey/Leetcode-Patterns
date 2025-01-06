package main

import "fmt"

// https://leetcode.com/problems/rotate-image/description/

// траснпонировали или обернули
func rotate(matrix [][]int)  {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < len(matrix[0]); j++ {
			temp := matrix[i][j];
            matrix[i][j] = matrix[j][i];
            matrix[j][i] = temp;
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < len(matrix[0])/2; j++ {
			temp := matrix[i][j]
            matrix[i][j] = matrix[i][n-j-1]
            matrix[i][n-j-1] = temp
		}
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
	printMatrix(matrix3x3)
	fmt.Println()

	rotate(matrix3x3)
	printMatrix(matrix3x3)
	fmt.Println()

	matrix4x4 := [][]int {
		{5,1,9,11},
		{2,4,8,10},
			{13,3,6,7},
				{15,14,12,16},
	}
	printMatrix(matrix4x4)
	fmt.Println()

	rotate(matrix4x4)
	printMatrix(matrix4x4)
	fmt.Println()

}