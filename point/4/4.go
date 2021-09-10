package main

import "fmt"

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(findNumberIn2DArray(matrix, 5))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	i, j := r-1, 0
	for i >= 0 && j < c {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--
		}
		if matrix[i][j] < target {
			j++
		}
	}
	return false
}
