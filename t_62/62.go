package main

import "fmt"


// начинаме с нижней точки и потом складываем правый и нижний путь
func uniquePaths(m int, n int) int {
	matrix := make([][]int, m) 
    for i := range matrix {
        matrix[i] = make([]int, n) 
    }

	matrix[m-1][n-1] = 1

	row := m - 1
	for row >= 0 {
		for i := n-1; i >= 0; i-- {
			if row + 1 < m && i + 1 < n {
				matrix[row][i] = matrix[row][i + 1] + matrix[row + 1][i]
			} else {
				matrix[row][i] = 1
			}
		}
		row--;
	}

	return matrix[0][0]
}

func main() {
	fmt.Println(uniquePaths(3,7))
}