package main

import (
	"fmt"
)

// https://leetcode.com/problems/word-search/

// это лучше и быстрее
// p.s. сама по себе задача несложная и интересная
// мб я визуализирую работу этого алгоритма
func traverse(board [][]byte, word string, i int, j int) bool {
	if len(word) == 0 {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[0] {
		return false
	}

	original_val := board[i][j]
	board[i][j] = '#'

	match := traverse(board, word[1:], i - 1, j) || traverse(board, word[1:], i + 1, j) || traverse(board, word[1:], i, j - 1) || traverse(board, word[1:], i, j + 1) 

	board[i][j] = original_val
	return match
}	

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && traverse(board, word, i, j) {
				return true
			}
        }
    }

	return false
}

// работает, но чисто визуально не нравится
// func traverse(board [][]byte, word string, currentX int, currentY int) bool {
// 	if len(word) == 0 {
// 		return true
// 	}

// 	if board[currentX][currentY] != word[0] {
// 		return false
// 	}

// 	temp := board[currentX][currentY]
// 	board[currentX][currentY] = '#'

// 	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

// 	for _, direction := range directions {
// 		newX, newY := currentX+direction[0], currentY+direction[1]

// 		if newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[0]) {
// 			if traverse(board, word[1:], newX, newY) {
// 				return true
// 			}
// 		}

// 	}
// 	board[currentX][currentY] = temp
// 	return false
// }	

// func exist(board [][]byte, word string) bool {
// 	if len(board) == 1 && len(word) == 1 && board[0][0] == word[0] {
// 		return true
// 	}

// 	for i := 0; i < len(board); i++ {
//         for j := 0; j < len(board[0]); j++ {
// 			if traverse(board, word, i, j) {
// 				return true
// 			}
//         }
//     }

// 	return false
// }


func main() {
    board := [][]byte{
        {'A', 'B', 'C', 'E'},
        {'S', 'F', 'C', 'S'},
        {'A', 'D', 'E', 'E'},
    }
    
    word := "ABCCED"
    
    found := exist(board, word)
    fmt.Println(found) 
}