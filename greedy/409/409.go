package main

import "fmt"

func main() {
	s := ""
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) int {
	count := 0
	flag := true
	slic := make([]int, 120)
	for i := 0; i < len(s); i++ {
		slic[s[i]] += 1
	}
	for j := 0; j < len(slic); j++ {
		count += (slic[j] / 2) * 2
		if flag && slic[j]%2 == 1 {
			count++
			flag = false
		}
	}
	return count
}
