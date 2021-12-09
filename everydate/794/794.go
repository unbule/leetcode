package main

import "fmt"

func main() {

}

func win(board []string, p byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == p && board[i][1] == p && board[i][2] == p ||
			board[0][i] == p && board[1][i] == p && board[2][i] == p {
			return true
		}
	}
	if board[0][0] == p && board[1][1] == p && board[2][2] == p ||
		board[0][2] == p && board[1][1] == p && board[2][0] == p {
		fmt.Println("123", string(p))
		return true
	}
	return false
}

func validTicTacToe(board []string) bool {
	countO := 0
	countX := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				countO++
			}
			if board[i][j] == 'X' {
				countX++
			}
		}
	}

	if win(board, 'X') && !win(board, 'O') && countX-1 == countO {
		fmt.Println("1112")
		return true
	}
	if win(board, 'O') && !win(board, 'X') && countX == countO {
		fmt.Println("111")
		return true
	}
	if !win(board, 'X') && !win(board, 'O') && (countX == countO || countX-1 == countO) {
		fmt.Println("1113")
		return true
	}
	return false
}
