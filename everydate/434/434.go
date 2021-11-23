package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ", , , ,        a, eaefa"
	fmt.Println(countSegments(s))
}

func countSegments(s string) int {
	s = strings.TrimSpace(s)
	// str := strings.Split(s, " ")
	// fmt.Println(str)
	// return len(str)
	p := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res += p
			p = 0
		} else {
			p = 1
		}
	}
	if p == 1 {
		res += p
	}
	return res
}
