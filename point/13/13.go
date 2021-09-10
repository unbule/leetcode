package main

import "fmt"

func main() {
	m, n := 16, 8
	fmt.Println(movingCount(m, n, 4))
}

func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	ans := 0
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i] = append(matrix[i], computer(i, j))
			//matrix[i][j] = computer(i, j)

		}
	}
	fmt.Println(matrix)
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if matrix[x][y] <= k {
				ans++
			}
		}
	}
	return ans
}

func computer(i, j int) int {
	ans := 0
	for i > 0 {
		ans = ans + i%10
		i = i / 10
	}
	for j > 0 {
		ans = ans + j%10
		j = j / 10
	}
	return ans
}
