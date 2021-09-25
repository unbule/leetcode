package main

import "fmt"

func main() {
	word1 := "sea"
	word2 := "eat"
	fmt.Println(minDistance(word1, word2))
}

//最长公共子序列
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, v1 := range word1 {
		for j, v2 := range word2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return m - dp[m][n] + n - dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
