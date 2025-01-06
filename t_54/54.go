package main

// https://leetcode.com/problems/spiral-matrix/

// духота, тупо не налажать в счетчиках 
func spiralOrder(matrix [][]int) []int {
    res := make([]int, 0)

	top := 0
	bottom := len(matrix) - 1

	left := 0
	right := len(matrix[0]) - 1
	direction := 0
	for top <= bottom && left <= right {
		if direction == 0 {
			
			for i := left; i < right + 1; i++ {
				res = append(res, matrix[top][i])
			}
			
			top++
			direction = 1
		} else if direction == 1 {

			for i := top; i < bottom + 1; i++ {
				res = append(res, matrix[i][right])
			}

			right--
			direction = 2
		} else if direction == 2 {

			for i := right; i > left - 1; i-- {
				res = append(res, matrix[bottom][i])
			}

			bottom--
			direction = 3
		} else if direction == 3{

			for i := bottom; i > top - 1; i--{
				res = append(res, matrix[i][left])
			}

			left++
			direction = 0
		}
	}

	return res
}