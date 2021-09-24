package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	var backtrace func([][]byte, int, int, string, int) bool
	backtrace = func(board [][]byte, x, y int, s string, index int) bool {
		if index == len(s) {
			return true
		}
		if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
			return false
		}
		if board[x][y] != s[index] {
			return false
		}
		board[x][y] ^= 255
		ex := backtrace(board, x+1, y, s, index+1) || backtrace(board, x-1, y, s, index+1) ||
			backtrace(board, x, y+1, s, index+1) || backtrace(board, x, y-1, s, index+1)
		board[x][y] ^= 255
		return ex
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrace(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}
