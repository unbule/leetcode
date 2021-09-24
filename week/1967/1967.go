package main

import (
	"fmt"
	"strings"
)

func main() {
	patterns := []string{"a", "abc", "bc", "d"}
	word := "abc"
	fmt.Println(numOfStrings(patterns, word))
}

func numOfStrings(patterns []string, word string) int {
	ans := 0
	for i := 0; i < len(patterns); i++ {
		if strings.Contains(word, patterns[i]) {
			ans++
		}
	}
	return ans
}
