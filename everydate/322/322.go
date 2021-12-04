package main

import (
	"fmt"
	"strings"
)

func main() {
	var str strings.Builder
	for i := 0; i < 1000; i++ {
		str.WriteString("a")
	}
	fmt.Println(str.String())
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}
			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
		fmt.Println(dp)
	}
	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}

			count := dp[i-c] + 1
			if dp[i] == -1 || count < dp[i] {
				dp[i] = count
			}
		}
		fmt.Println(dp)
	}
	return dp[amount]
}
