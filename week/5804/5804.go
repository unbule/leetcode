package main

import "fmt"

func main() {
	s := "aabbccd"
	fmt.Println(areOccurrencesEqual(s))
}

func areOccurrencesEqual(s string) bool {
	slic := make([]int, 27)
	for i := 0; i < len(s); i++ {
		slic[int(s[i]-'a')]++
	}
	target := 0
	for _, c := range slic {
		if c == 0 {
			continue
		}
		if target == 0 {
			target = c
		} else if c != target {
			return false
		}
	}
	return true
}
