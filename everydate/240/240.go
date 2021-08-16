package main

import "fmt"

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 5))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)    //行
	n := len(matrix[0]) //列
	i, j := m-1, 0
	for i >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		}
	}
	return false
}
