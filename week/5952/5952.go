package main

import (
	"fmt"
	"strings"
)

func main() {
	countPoints("B0B6G0R6R0R6G9")
}

func countPoints(rings string) int {
	rem := make(map[int]string)
	res := 0
	for i := 0; i < 10; i++ {
		rem[i] = "RBG"
	}
	for i := 0; i < len(rings)-1; i = i + 2 {
		j := i + 1
		index := rings[j] - '0'
		rem[int(index)] = strings.ReplaceAll(rem[int(index)], string(rings[i]), "")
	}
	fmt.Println(rem)
	for _, v := range rem {
		if v == "" {
			res++
		}
	}
	return res
}
