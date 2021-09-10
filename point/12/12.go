package main

func main() {

}

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	var backtrace func([][]byte, int, int, int) bool
	backtrace = func(board [][]byte, r, c, index int) bool {
		if board[r][c] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		temp := board[r][c]
		board[r][c] = byte(' ')
		if r-1 >= 0 && backtrace(board, r-1, c, index+1) {
			return true
		}
		if r+1 < row && backtrace(board, r+1, c, index+1) {
			return true
		}
		if c-1 >= 0 && backtrace(board, r, c-1, index+1) {
			return true
		}
		if c+1 < col && backtrace(board, r, c+1, index+1) {
			return true
		}
		board[r][c] = temp
		return false
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if backtrace(board, i, j, 0) {
				return true
			}
		}
	}
	return false
}
