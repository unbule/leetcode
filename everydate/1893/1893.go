package main

import "fmt"

func main() {
	s := [][]int{{1, 2}, {3, 4}, {5, 6}}
	if isCovered(s, 2, 5) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}

func isCovered(ranges [][]int, left int, right int) bool {
	chafen := make([]int, 52)
	for _, n := range ranges {
		chafen[n[0]]++
		chafen[n[1]+1]--
	}

	sum := 0
	for i := 1; i < len(chafen); i++ {
		sum = sum + chafen[i]
		if sum <= 0 && left <= i && i <= right {
			return false
		}
	}
	return true
}
