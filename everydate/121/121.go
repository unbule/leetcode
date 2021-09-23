package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	min := math.MaxInt64
	max := 0
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if max < prices[i]-min {
			max = prices[i] - min
		}
	}
	return max
}
